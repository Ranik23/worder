package grpcserver

import (
	"Worder/internal/api/gen"
	corrector_service "Worder/internal/services/corrector"
	"context"
	"google.golang.org/grpc"
)

type CorrectWordServer struct {
	gen.UnimplementedWordCorrectorServer
	correctService corrector_service.Corrector
}

func Register(gRPC *grpc.Server, correctService corrector_service.Corrector) {
	gen.RegisterWordCorrectorServer(gRPC, &CorrectWordServer{correctService: correctService})
}


func (s *CorrectWordServer) CorrectWord(ctx context.Context, word *gen.CorrectWordRequest) (*gen.CorrectWordResponse, error) {
	new_word, ok := s.correctService.CorrectWord(word.String())
	if !ok {
		new_word = "unrecognized"
	}
	return &gen.CorrectWordResponse{
		CorrectedWord: new_word,
	}, nil
}