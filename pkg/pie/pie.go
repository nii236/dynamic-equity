package pie

import (
	"time"

	"github.com/zpatrick/rbac"
)

// Pie holds all dynamic equity data without calcs
type Pie struct {
	Roles                 []*rbac.Role
	Currency              int
	NonCashMultiplier     int
	CashMultiplier        int
	CommissionRatePercent int
	RoyaltyRatePercent    int
	FindersFeePercent     int
	Projects              []*Project
	Chunks                []*Chunk
}

// NewPie returns a new pie
func NewPie(
	Roles []*rbac.Role,
	Currency int,
	NonCashMultiplier int,
	CashMultiplier int,
	CommissionRatePercent int,
	RoyaltyRatePercent int,
	FindersFeePercent int,
) *Pie {
	return &Pie{
		Roles,
		Currency,
		NonCashMultiplier,
		CashMultiplier,
		CommissionRatePercent,
		RoyaltyRatePercent,
		FindersFeePercent,
		[]*Project{},
		[]*Chunk{},
	}
}

// Chunk is a collection of slices with metadata
type Chunk struct {
	Amount    int
	User      string
	CreatedAt time.Time
}

// AddChunk to a pie
func (p *Pie) AddChunk(amount int, user string) {
	p.Chunks = append(p.Chunks, &Chunk{amount, user, time.Now()})
}

// AddProject to a pie
func (p *Pie) AddProject(name string) {
	p.Projects = append(p.Projects, &Project{name})
}

type Project struct {
	Name string
}
