package api

import (
	"context"
	pb "dynamic-equity/proto/dynamicequity"
)

// PieController holds pie methods
type PieController struct {
}

// NewPieController returns a new controller
func NewPieController() *PieController {
	return &PieController{}
}

// List pies
func (c *PieController) List(context.Context, *pb.PieListRequest) (*pb.PieListResponse, error) {
	return nil, nil
}

// Create pies
func (c *PieController) Create(context.Context, *pb.PieCreateRequest) (*pb.PieCreateResponse, error) {
	return nil, nil
}

// Archive pies
func (c *PieController) Archive(context.Context, *pb.PieArchiveRequest) (*pb.PieArchiveResponse, error) {
	return nil, nil
}

// ProjectList for existing pie
func (c *PieController) ProjectList(context.Context, *pb.PieProjectListRequest) (*pb.PieProjectListResponse, error) {
	return nil, nil
}

// ProjectAdd for existing pie
func (c *PieController) ProjectAdd(context.Context, *pb.PieProjectAddRequest) (*pb.PieProjectAddResponse, error) {
	return nil, nil
}

// ProjectArchive for existing pie
func (c *PieController) ProjectArchive(context.Context, *pb.PieProjectArchiveRequest) (*pb.PieProjectArchiveResponse, error) {
	return nil, nil
}

// UserList for existing pie
func (c *PieController) UserList(context.Context, *pb.PieUserListRequest) (*pb.PieUserListResponse, error) {
	return nil, nil
}

// UserAdd for existing pie
func (c *PieController) UserAdd(context.Context, *pb.PieUserAddRequest) (*pb.PieUserAddResponse, error) {
	return nil, nil
}

// UserRemove for existing pie
func (c *PieController) UserRemove(context.Context, *pb.PieUserRemoveRequest) (*pb.PieUserRemoveResponse, error) {
	return nil, nil
}

// ChunkList for existing pie
func (c *PieController) ChunkList(context.Context, *pb.PieChunkListRequest) (*pb.PieChunkListResponse, error) {
	return nil, nil
}

// ChunkCreate for existing pie
func (c *PieController) ChunkCreate(context.Context, *pb.PieChunkCreateRequest) (*pb.PieChunkCreateResponse, error) {
	return nil, nil
}

// ChunkApprove for existing pie
func (c *PieController) ChunkApprove(context.Context, *pb.PieChunkApproveRequest) (*pb.PieChunkApproveResponse, error) {
	return nil, nil
}

// ChunkReject for existing pie
func (c *PieController) ChunkReject(context.Context, *pb.PieChunkRejectRequest) (*pb.PieChunkRejectResponse, error) {
	return nil, nil
}
