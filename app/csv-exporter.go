package app

import (
	"bufio"
	"encoding/csv"
	"os"
	"strings"

	uuid "github.com/satori/go.uuid"
)

const CsvDateTimeFormat = "2006-01-02 15:04:05" // nolint: golint

const CsvDateFormat = "2006-01-02" // nolint: golint

func dashlessUUID(id uuid.UUID) string {
	return strings.Replace(id.String(), "-", "", -1)
}

func exportTx(writer *csv.Writer, tx *FlattenedTransaction) {
	err := writer.Write([]string{
		tx.Date.Format(CsvDateFormat),
		tx.Description,
		tx.Account1,
		tx.AccType1,
		tx.Value1.FloatString(2),
		tx.Account2,
		tx.AccType2,
		tx.Value2.FloatString(2),
		dashlessUUID(tx.Id),
		tx.Creation.Format(CsvDateTimeFormat),
		dashlessUUID(tx.Ref1),
		dashlessUUID(tx.Ref2),
	})
	if err != nil {
		panic(err)
	}
}

func exportBalance(writer *csv.Writer, balance *AccountBalance) {
	err := writer.Write([]string{
		balance.Name,
		balance.Date.Format(CsvDateFormat),
		balance.Balance.FloatString(2),
		dashlessUUID(balance.Transaction),
		balance.Type,
		dashlessUUID(balance.Id),
		dashlessUUID(balance.Parent),
	})
	if err != nil {
		panic(err)
	}
}

// ExportTransactions - writes flattened transactions to a csv file
func ExportTransactions(txs []FlattenedTransaction, fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer func() {
		err2 := file.Close()
		if err2 != nil {
			panic(err2)
		}
	}()
	writer := csv.NewWriter(bufio.NewWriter(file))
	defer writer.Flush()
	err = writer.Write([]string{"Date", "Description",
		"Account1", "AccType1", "Value1",
		"Account2", "AccType2", "Value2",
		"Id", "Creation", "Ref1", "Ref2"})
	if err != nil {
		panic(err)
	}
	for _, tx := range txs {
		exportTx(writer, &tx)
	}
}

// ExportBalances - writes account balances to a csv file
func ExportBalances(balances []AccountBalance, fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer func() {
		err2 := file.Close()
		if err2 != nil {
			panic(err2)
		}
	}()
	writer := csv.NewWriter(bufio.NewWriter(file))
	defer writer.Flush()
	err = writer.Write([]string{
		"Name", "Date", "Balance", "Transaction", "Type", "Id", "Parent"})
	if err != nil {
		panic(err)
	}
	for _, balance := range balances {
		exportBalance(writer, &balance)
	}
}
