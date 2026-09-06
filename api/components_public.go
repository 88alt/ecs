//go:build ecs_public

package api

import (
	"context"
	"fmt"
	"time"

	speedmodel "github.com/oneclickvirt/speedtest/model"
)

func hasPrivateComponentData() bool { return false }

func loadSecurityComponentData(ctx context.Context, _ bool) componentDataResult {
	err := fmt.Errorf("security component unavailable in public build")
	return failedComponentData(ctx, dnsblDataFile, err)
}

func loadPrivateSpeedComponentData(ctx context.Context, _ bool) componentDataResult {
	err := fmt.Errorf("private speed component unavailable in public build")
	return failedComponentData(ctx, privateDataFile, err)
}

func loadPrivateSpeedComponentDataWithNetwork(ctx context.Context, offline bool, _ speedmodel.Network) componentDataResult {
	return loadPrivateSpeedComponentData(ctx, offline)
}

func loadTransferComponentData(ctx context.Context, _ bool) componentDataResult {
	err := fmt.Errorf("transfer component unavailable in public build")
	return failedComponentData(ctx, transferDataFile, err)
}

func collectSecurityComponent(context.Context, string, string, []dnsblZoneInput) ComponentReport {
	report := componentPayload("security.evidence", "goecs.security/v1", ReportStatusUnavailable, time.Now(), nil, nil)
	report.Reason = "security component unavailable in public build"
	return report
}

func runPrivateSpeedBenchmarks(context.Context, int) (any, int, []privateSpeedBenchmark) {
	return nil, 0, nil
}

func runEmbeddedPrivateSpeedBenchmarks(context.Context, int) (any, int, []privateSpeedBenchmark) {
	return nil, 0, nil
}

func runInternationalPrivateSpeedBenchmarks(context.Context, int) (any, int, []privateSpeedBenchmark) {
	return nil, 0, nil
}

func runEmbeddedInternationalPrivateSpeedBenchmarks(context.Context, int) (any, int, []privateSpeedBenchmark) {
	return nil, 0, nil
}
