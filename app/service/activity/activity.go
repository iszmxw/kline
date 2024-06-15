package activity

import "goapi/app/models"

func GetBonusAmounts(rankNo int, tSystemRankRewards []models.TSystemRankReward) float64 {
	bonusAmounts := 0.0
	for _, rankReward := range tSystemRankRewards {
		if rankReward.MinRank == 0 {
			if int(rankReward.MaxRank) == rankNo {
				bonusAmounts = rankReward.RewardAmount
			}
		} else {
			if rankNo >= int(rankReward.MaxRank) && rankNo <= int(rankReward.MinRank) {
				bonusAmounts = rankReward.RewardAmount
			}
		}
	}
	return bonusAmounts
}
