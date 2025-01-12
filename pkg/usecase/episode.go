package usecase

import (
	"context"

	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/entity"
	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/errcode"
	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/repository"
)

type ListEpisodesRequest struct {
	Limit    int
	Offset   int
	SeasonID string
	SeriesID string
}

type ListEpisodesResponse struct {
	Episodes entity.Episodes
	Series   entity.SeriesMulti
	Seasons  entity.Seasons
}

func (u *UsecaseImpl) ListEpisodes(ctx context.Context, req *ListEpisodesRequest) (*ListEpisodesResponse, error) {
	ctx, span := tracer.Start(ctx, "usecase.UsecaseImpl#ListEpisodes")
	defer span.End()

	params := &repository.ListEpisodesParams{
		Limit:    req.Limit,
		Offset:   req.Offset,
		SeasonID: req.SeasonID,
	}
	episodes, err := u.db.Episode.List(ctx, params)
	if err != nil {
		return nil, errcode.New(err)
	}

	series := make(entity.SeriesMulti, 0, len(episodes))
	batchSeries, err := u.db.Series.BatchGet(ctx, series.SeriesIDs())
	if err != nil {
		return nil, errcode.New(err)
	}

	seasons := make(entity.Seasons, 0, len(episodes))
	batchSeason, err := u.db.Season.BatchGet(ctx, seasons.SeasonIDs())
	if err != nil {
		return nil, errcode.New(err)
	}

	return &ListEpisodesResponse{
		Episodes: episodes,
		Series:   batchSeries,
		Seasons:  batchSeason,
	}, nil
}
