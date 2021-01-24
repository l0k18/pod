package ffldb

import (
	"os"
	"path/filepath"
	"testing"

	chaincfg "github.com/l0k18/pod/pkg/chain/config"
	database "github.com/l0k18/pod/pkg/db"
	"github.com/l0k18/pod/pkg/util"
)

// BenchmarkBlockHeader benchmarks how long it takes to load the mainnet genesis block header.
func BenchmarkBlockHeader(b *testing.B) {
	// Start by creating a new database and populating it with the mainnet genesis block.
	dbPath := filepath.Join(os.TempDir(), "ffldb-benchblkhdr")
	_ = os.RemoveAll(dbPath)
	db, err := database.Create("ffldb", dbPath, blockDataNet)
	if err != nil {
		b.Fatal(err)
	}
	defer func() {
		if err := os.RemoveAll(dbPath); Check(err) {
		}
	}()
	defer func() {
		if err := db.Close(); Check(err) {
		}
	}()
	err = db.Update(func(tx database.Tx) error {
		block := util.NewBlock(chaincfg.MainNetParams.GenesisBlock)
		return tx.StoreBlock(block)
	})
	if err != nil {
		b.Fatal(err)
	}
	b.ReportAllocs()
	b.ResetTimer()
	err = db.View(func(tx database.Tx) error {
		blockHash := chaincfg.MainNetParams.GenesisHash
		for i := 0; i < b.N; i++ {
			_, err := tx.FetchBlockHeader(blockHash)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		b.Fatal(err)
	}
	// Don't benchmark teardown.
	b.StopTimer()
}

// BenchmarkBlockHeader benchmarks how long it takes to load the mainnet genesis block.
func BenchmarkBlock(b *testing.B) {
	// Start by creating a new database and populating it with the mainnet genesis block.
	dbPath := filepath.Join(os.TempDir(), "ffldb-benchblk")
	_ = os.RemoveAll(dbPath)
	db, err := database.Create("ffldb", dbPath, blockDataNet)
	if err != nil {
		b.Fatal(err)
	}
	defer func() {
		if err := os.RemoveAll(dbPath); Check(err) {
		}
	}()
	defer func() {
		if err := db.Close(); Check(err) {
		}
	}()
	err = db.Update(func(tx database.Tx) error {
		block := util.NewBlock(chaincfg.MainNetParams.GenesisBlock)
		return tx.StoreBlock(block)
	})
	if err != nil {
		b.Fatal(err)
	}
	b.ReportAllocs()
	b.ResetTimer()
	err = db.View(func(tx database.Tx) error {
		blockHash := chaincfg.MainNetParams.GenesisHash
		for i := 0; i < b.N; i++ {
			_, err := tx.FetchBlock(blockHash)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		b.Fatal(err)
	}
	// Don't benchmark teardown.
	b.StopTimer()
}
