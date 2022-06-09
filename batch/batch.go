package batch

import (
	"context"
	"dc-stats/batch/tasks"
	"dc-stats/constants"
	"log"
	"sync"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/procyon-projects/chrono"
)

var (
	fetchVoiceDataTask     *chrono.ScheduledTask
	processVoiceDataTask   *chrono.ScheduledTask
	fetchMessageDataTask   *chrono.ScheduledTask
	processMessageDataTask *chrono.ScheduledTask
	fetchNicknamesTask     *chrono.ScheduledTask
	fetchChannelNamesTask  *chrono.ScheduledTask
)

func Start(goBot *discordgo.Session) {

	taskScheduler := chrono.NewDefaultTaskScheduler()

	/**
	 * 	GATHER VOICE STATS TASK
	 */
	task, err := taskScheduler.ScheduleWithFixedDelay(func(ctx context.Context) {
		var wg sync.WaitGroup
		wg.Add(1)
		go tasks.GatherVoiceStats(goBot, ctx, &wg)
		wg.Wait()
	}, constants.FetchVoiceDataInterval, chrono.WithTime(<-time.After(constants.FetchVoiceDataInterval)))
	fetchVoiceDataTask = &task

	if err == nil {
		log.Print("FetchVoiceDataTask has been scheduled successfully.  Fixed delay: ", constants.FetchVoiceDataInterval)
	}

	/**
	 * 	PROCESS VOICE STATS TASK
	 */
	task, err = taskScheduler.ScheduleWithFixedDelay(func(ctx context.Context) {
		var wg sync.WaitGroup
		wg.Add(1)
		go tasks.ProcessVoiceStats(goBot, ctx, &wg)
		wg.Wait()
	}, constants.ProcessVoiceDataInterval, chrono.WithTime(<-time.After(constants.ProcessVoiceDataInterval)))
	processVoiceDataTask = &task

	if err == nil {
		log.Print("ProcessVoiceDataTask has been scheduled successfully. Fixed delay: ", constants.ProcessVoiceDataInterval)
	}

	/**
	 *  GATHER MESSAGE STATS TASK
	 */
	task, err = taskScheduler.ScheduleWithFixedDelay(func(ctx context.Context) {
		var wg sync.WaitGroup
		wg.Add(1)
		go tasks.GatherMessageStats(goBot, ctx, &wg)
		wg.Wait()
	}, constants.FetchMessageDataInterval, chrono.WithTime(<-time.After(constants.FetchMessageDataInterval)))
	fetchMessageDataTask = &task

	if err == nil {
		log.Print("FetchMessageDataTask has been scheduled successfully. Fixed delay: ", constants.FetchMessageDataInterval)
	}

	/**
	 * 	PROCESS MESSAGE STATS TASK
	 */
	task, err = taskScheduler.ScheduleWithFixedDelay(func(ctx context.Context) {
		var wg sync.WaitGroup
		wg.Add(1)
		go tasks.ProcessMessageStats(goBot, ctx, &wg)
		wg.Wait()
	}, constants.ProcessMessageDataInterval, chrono.WithTime(<-time.After(constants.ProcessMessageDataInterval)))
	processMessageDataTask = &task

	if err == nil {
		log.Print("ProcessMessageDataTask has been scheduled successfully. Fixed delay: ", constants.ProcessMessageDataInterval)
	}

	/**
	 * 	GATHER USER NICKNAMES
	 */
	task, err = taskScheduler.ScheduleWithFixedDelay(func(ctx context.Context) {
		var wg sync.WaitGroup
		wg.Add(1)
		go tasks.GatherNicknameStats(goBot, ctx, &wg)
		wg.Wait()
	}, constants.FetchNicknamesInterval, chrono.WithTime(time.Now()))
	fetchNicknamesTask = &task

	if err == nil {
		log.Print("FetchNicknamesTask has been scheduled successfully.  Fixed delay: ", constants.FetchNicknamesInterval)
	}

	/**
	 * 	GATHER CHANNEL NAMES
	 */
	task, err = taskScheduler.ScheduleWithFixedDelay(func(ctx context.Context) {
		var wg sync.WaitGroup
		wg.Add(1)
		go tasks.GatherChannelNameStats(goBot, ctx, &wg)
		wg.Wait()
	}, constants.FetchChannelNamesInterval, chrono.WithTime(time.Now()))
	fetchChannelNamesTask = &task

	if err == nil {
		log.Print("FetchChannelNamesTask has been scheduled successfully.  Fixed delay: ", constants.FetchChannelNamesInterval)
	}

}
