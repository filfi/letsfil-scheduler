package glif

import (
	"bytes"
	"context"
	"fmt"
	"git.sxxfuture.net/filfi/letsfil/letsfil-job/src/pkg/log"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/lotus/api"
	//"github.com/filecoin-project/lotus/chain/actors/builtin/miner"
	"github.com/filecoin-project/go-state-types/builtin/v9/miner"
	chaintypes "github.com/filecoin-project/lotus/chain/types"

	"golang.org/x/xerrors"
)

func GetMinerPower(miner string) (mbi *api.MinerPower, err error) {

	mineraddr, err := address.NewFromString(miner)
	fmt.Printf("mineraddr: %s\n", mineraddr)
	if err != nil {
		return nil, err
	}

	//currHead, err := lapi.ChainHead(ctx)

	mbi, err = lapi.StateMinerPower(ctx, mineraddr, chaintypes.EmptyTSK)
	if err != nil {
		log.Errorf("failed to get mining base info: %s", err)
		err = xerrors.Errorf("failed to get mining base info: %w", err)
		return nil, err
	}
	return mbi, nil

}

func GetMinerPowerV1(minerId string) (mbi *api.MiningBaseInfo, err error) {

	mineraddr, err := address.NewFromString(minerId)
	fmt.Printf("mineraddr: %s\n", mineraddr)
	if err != nil {
		return nil, err
	}
	currHead, err := lapi.ChainHead(ctx)
	//currEpoch := currHead.Height()

	tsk := chaintypes.NewTipSetKey(currHead.Cids()...)

	ma, err := lapi.StateGetActor(context.Background(), mineraddr, tsk)
	if err != nil {
		return nil, err
	}

	//actorHead, err := service.RequestChainReadObj(ma.Head.String())
	actorHead, err := lapi.ChainReadObj(ctx, ma.Head)
	if err != nil {
		return nil, err
	}

	var minerState miner.State

	if err := minerState.UnmarshalCBOR(bytes.NewReader(actorHead)); err != nil {
		return nil, err
	}

	actorInfo, err := lapi.ChainReadObj(ctx, minerState.Info)
	if err != nil {
		return nil, err
	}
	var minerInfo miner.MinerInfo

	if err := minerInfo.UnmarshalCBOR(bytes.NewReader(actorInfo)); err != nil {
		return nil, err
	}

	return nil, nil

}

func GetMinerState(minerId string) (mbs *miner.State, err error) {

	mineraddr, err := address.NewFromString(minerId)
	fmt.Printf("mineraddr: %s\n", mineraddr)
	if err != nil {
		return nil, err
	}

	currHead, err := lapi.ChainHead(ctx)
	//currEpoch := currHead.Height()

	tsk := chaintypes.NewTipSetKey(currHead.Cids()...)

	ma, err := lapi.StateGetActor(context.Background(), mineraddr, tsk)
	if err != nil {
		return nil, err
	}

	//actorHead, err := service.RequestChainReadObj(ma.Head.String())
	actorHead, err := lapi.ChainReadObj(ctx, ma.Head)
	if err != nil {
		return nil, err
	}

	var minerState miner.State

	if err := minerState.UnmarshalCBOR(bytes.NewReader(actorHead)); err != nil {
		return nil, err
	}

	return &minerState, nil

}

func GetMinerSectors(minerId string) (mbi *api.MinerSectors, err error) {

	log.Debug("GetMinerSectors start")
	log.Infof("minerId: %s\n", minerId)
	mineraddr, err := address.NewFromString(minerId)
	log.Infof("mineraddr: %s\n", mineraddr.String())
	if err != nil {
		return nil, err
	}

	currHead, err := lapi.ChainHead(ctx)
	//currEpoch := currHead.Height()

	tsk := chaintypes.NewTipSetKey(currHead.Cids()...)

	ma, err := lapi.StateGetActor(context.Background(), mineraddr, tsk)
	if err != nil {
		return nil, err
	}

	//actorHead, err := service.RequestChainReadObj(ma.Head.String())
	actorHead, err := lapi.ChainReadObj(ctx, ma.Head)
	if err != nil {
		return nil, err
	}

	var minerState miner.State

	if err := minerState.UnmarshalCBOR(bytes.NewReader(actorHead)); err != nil {
		return nil, err
	}

	actorSectors, err := lapi.StateMinerSectorCount(ctx, mineraddr, tsk)
	if err != nil {
		return nil, err
	}
	log.Infof("actorSectors: %+v", actorSectors)

	return &actorSectors, nil

}

func GetChainHead() (mbs *chaintypes.TipSet, err error) {

	currHead, err := lapi.ChainHead(ctx)

	return currHead, nil

}
