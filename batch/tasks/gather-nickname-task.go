package tasks

import (
	"context"
	"dc-stats/database"
	"dc-stats/utils"
	"log"
	"sync"

	"github.com/bwmarrin/discordgo"
)

func GatherNicknameStats(goBot *discordgo.Session, ctx context.Context, wait *sync.WaitGroup) {
	defer wait.Done()

	for _, guild := range goBot.State.Guilds {
		guildId := guild.ID
		guildObject, err := database.FindOrCreateDataGuild(guildId, ctx)
		if err != nil {
			log.Println("Error finding/creating guild", err)
			continue
		}
		lastId := ""
		for {
			members, err := goBot.GuildMembers(guildId, lastId, 1000)
			if err != nil || len(members) == 0 {
				break
			}
			lastId = members[len(members)-1].User.ID
			for _, member := range members {
				userId := member.User.ID
				nickName := utils.GetUserNickName(*member)
				guildObject.UserNicknameMap[userId] = nickName
			}
			database.UpdateDataGuildUserNicknameMap(guildObject, ctx)
		}
	}

}
