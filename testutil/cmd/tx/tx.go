package tx

import (
	"context"
	"errors"
	"regexp"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/tokenize-titan/titan/testutil"
	"github.com/tokenize-titan/titan/testutil/cmd"
)

var rpcErrPattern = regexp.MustCompile(`RPC\serror\s(-?[\d]+)`)

type Tx struct {
	Height    testutil.Int    `json:"height"`
	Code      int             `json:"code"`
	Hash      string          `json:"txhash"`
	RawLog    string          `json:"raw_log"`
	GasUsed   testutil.BigInt `json:"gas_used"`
	GasWanted testutil.BigInt `json:"gas_wanted"`
	Events    []Event         `json:"events"`
}

type Event struct {
	Type       string      `json:"type"`
	Attributes []Attribute `json:"attributes"`
}

type Attribute struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	Index bool   `json:"index"`
}

func QueryTx(ctx context.Context, txHash string) (*Tx, error) {
	for {
		output, err := cmd.Exec("titand", "query", "tx", txHash, "--output=json")
		if err != nil {
			matches := rpcErrPattern.FindStringSubmatch(string(output))
			if len(matches) == 2 && matches[1] == "-32603" {
				// Transaction not found, wait until it is delivered or timeout
				select {
				case <-time.After(1 * time.Second):
					continue
				case <-ctx.Done():
					if err == nil {
						err = ctx.Err()
					}
					return nil, err
				}
			}
			return nil, err
		}
		var tx Tx
		if err := cmd.UnmarshalJSON(output, &tx); err != nil {
			return nil, err
		}
		return &tx, nil
	}
}

func ExecTx(ctx context.Context, args ...string) (*Tx, error) {
	args = append([]string{"tx"}, args...)
	args = append(args, "--gas=auto", "--gas-adjustment=2", "--gas-prices=10utkx", "--output=json", "-y")
	output, err := cmd.Exec("titand", args...)
	if err != nil {
		return nil, err
	}
	var tx Tx
	if err := cmd.UnmarshalJSON(output, &tx); err != nil {
		return nil, err
	}
	if tx.Code != 0 {
		return nil, errors.New(tx.RawLog)
	}
	return QueryTx(ctx, tx.Hash)
}

func MustExecTx(t testing.TB, ctx context.Context, args ...string) Tx {
	tx, err := ExecTx(ctx, args...)
	require.NoError(t, err)
	require.Equal(t, 0, tx.Code, tx.RawLog)
	return *tx
}