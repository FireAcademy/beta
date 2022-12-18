// credits: https://sql2gorm.mccode.info/
package main


type SyncedBlock struct {
	HeaderHash string `gorm:"column:header_hash;primary_key"`
	Height     int64  `gorm:"column:height;NOT NULL"`
}


type SingletonState struct {
	CoinID          string `gorm:"column:coin_id;primary_key"`
	Melted          int    `gorm:"column:melted;NOT NULL"`

	HeaderHash      string `gorm:"column:header_hash;NOT NULL"`
	Height          int64  `gorm:"column:height;NOT NULL"`

	ParentCoinID    string `gorm:"column:parent_coin_id;NOT NULL"`
	PuzzleHash      string `gorm:"column:puzzle_hash;NOT NULL"`
	Amount          int64  `gorm:"column:amount;NOT NULL"`

	LauncherID      string `gorm:"column:launcher_id;NOT NULL"`
	InnerPuzzleHash string `gorm:"column:inner_puzzle_hash"`
}


type Puzzle struct {
	PuzzleHash string `gorm:"column:puzzle_hash;primary_key"`
	Puzzle     []byte `gorm:"column:puzzle;type:bytea;NOT NULL"`
}

