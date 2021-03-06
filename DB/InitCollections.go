package DB

import (
	"github.com/SommerEngineering/Ocean/CustomerDB"
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
)

//initCollections initializes collection and ensures presencs of index keys.
func initCollections() {

	// Get the database:
	dbSession, db := CustomerDB.DB()
	defer dbSession.Close()

	if db == nil { // Database not found
		Log.LogFull(senderName, LM.CategoryAPP, LM.LevelERROR, LM.SeverityCritical, LM.ImpactCritical, LM.MessageNameDATABASE, `Was not able to get the customer database.`)
		return
	}

	// Get all collections:
	ocollAnswers := db.C(collAnswers)
	ocollRecommendations := db.C(collRecommendations)
	ocollFeedback := db.C(collFeedback)

	// Ensure the indexes:
	ocollAnswers.EnsureIndexKey(`Version`)
	ocollAnswers.EnsureIndexKey(`Session`)
	ocollAnswers.EnsureIndexKey(`TimeUTC`)
	ocollAnswers.EnsureIndexKey(`Version`, `TimeUTC`)

	ocollRecommendations.EnsureIndexKey(`Version`)
	ocollRecommendations.EnsureIndexKey(`Session`)
	ocollRecommendations.EnsureIndexKey(`CreateTimeUTC`)
	ocollRecommendations.EnsureIndexKey(`Version`, `CreateTimeUTC`)

	ocollFeedback.EnsureIndexKey(`Session`)
	ocollFeedback.EnsureIndexKey(`Rating`)
}
