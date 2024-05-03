package main

import (
	"context"
	"fmt"

	"group_service/cmd/model"
	u "group_service/cmd/utils"
	pb "group_service/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

// AddStudent, CreateGroup, DeleteGroup, FindGroup, FindGroupsTeacher, FindGroupsStudent, RemoveStudent

func (s *Server) AddStudent(ctx context.Context, in *pb.AddStudentRequest) (*emptypb.Empty, error) {
	userId, err := u.GetUserIdFromToken(in.Token)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}

	err = s.GroupRepository.AddStudent(ctx, userId, in.GroupId)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *Server) CreateGroup(ctx context.Context, in *pb.CreateGroupRequest) (*emptypb.Empty, error) {
	userId, err := u.GetUserIdFromToken(in.Token)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}

	newGroup := model.Group{
		Title:     in.Title,
		TeacherId: userId,
	}

	err = s.GroupRepository.CreateGroup(ctx, newGroup)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}

	return &emptypb.Empty{}, nil
}

func (s *Server) DeleteGroup(ctx context.Context, in *pb.DeleteGroupRequest) (*emptypb.Empty, error) {
	userId, err := u.GetUserIdFromToken(in.Token)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}

	if IsTeacher := s.GroupRepository.IsTeacher(ctx, userId, in.GroupId); IsTeacher {
		err = s.GroupRepository.DeleteGroup(ctx, in.GroupId)
		if err != nil {
			return nil, status.Errorf(
				codes.Internal,
				fmt.Sprintf("Internal error: %v", err),
			)
		}
	} else {
		return nil, status.Errorf(
			codes.PermissionDenied,
			"You are not allowed to delete the group",
		)
	}

	return &emptypb.Empty{}, nil
}

func (s *Server) FindGroup(ctx context.Context, in *pb.FindGroupRequest) (*pb.GroupResponse, error) {
	userId, err := u.GetUserIdFromToken(in.Token)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}

	IsStudent := s.GroupRepository.IsStudent(ctx, userId, in.GroupId)
	IsTeacher := s.GroupRepository.IsTeacher(ctx, userId, in.GroupId)

	if IsStudent || IsTeacher {
		group, err := s.GroupRepository.FindById(ctx, in.GroupId)
		if err != nil {
			return nil, status.Errorf(
				codes.Internal,
				fmt.Sprintf("Internal error: %v", err),
			)
		}

		gr := &pb.GroupResponse{
			GroupId:  in.GroupId,
			Title:    group.Title,
			Students: group.Students,
		}

		return gr, nil
	} else {
		return nil, status.Errorf(
			codes.PermissionDenied,
			"You are not in the group",
		)
	}
}

func (s *Server) FindGroupsTeacher(in *pb.FindGroupsTeacherRequest, stream pb.GroupService_FindGroupsTeacherServer) error {
	userId, err := u.GetUserIdFromToken(in.Token)
	if err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}

	teacher_groups, teacher_groups_err := s.GroupRepository.FindByTeacherId(context.Background(), userId)
	if teacher_groups_err != nil {
		return teacher_groups_err
	}

	for _, teacher_group := range teacher_groups {
		stream_err := stream.Send(&pb.GroupResponse{
			GroupId:  teacher_group.Id.Hex(),
			Title:    teacher_group.Title,
			Students: teacher_group.Students,
		})
		if stream_err != nil {
			return stream_err
		}
	}

	return nil
}

func (s *Server) FindGroupsStudent(in *pb.FindGroupsStudentRequest, stream pb.GroupService_FindGroupsStudentServer) error {
	userId, err := u.GetUserIdFromToken(in.Token)
	if err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}

	groups, groups_err := s.GroupRepository.FindByStudentId(context.Background(), userId)
	if groups_err != nil {
		return groups_err
	}

	for _, group := range groups {
		stream_err := stream.Send(&pb.GroupResponse{
			GroupId:  group.Id.Hex(),
			Title:    group.Title,
			Students: group.Students,
		})
		if stream_err != nil {
			return stream_err
		}
	}

	return nil
}

func (s *Server) RemoveStudent(ctx context.Context, in *pb.RemoveStudentRequest) (*emptypb.Empty, error) {
	userId, err := u.GetUserIdFromToken(in.Token)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}

	if IsStudent := s.GroupRepository.IsStudent(ctx, userId, in.GroupId); IsStudent {
		err = s.GroupRepository.RemoveStudent(ctx, userId, in.GroupId)
		if err != nil {
			return nil, status.Errorf(
				codes.Internal,
				fmt.Sprintf("Internal error: %v", err),
			)
		}
	} else {
		return nil, status.Errorf(
			codes.PermissionDenied,
			"You are not allowed to remove the student",
		)
	}

	return &emptypb.Empty{}, nil
}
