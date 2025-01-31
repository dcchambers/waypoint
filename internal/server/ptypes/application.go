package ptypes

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/imdario/mergo"
	"github.com/mitchellh/go-testing-interface"
	"github.com/stretchr/testify/require"

	"github.com/hashicorp/waypoint/internal/pkg/validationext"
	pb "github.com/hashicorp/waypoint/internal/server/gen"
)

// TestApplication returns a valid project for tests.
func TestApplication(t testing.T, src *pb.Application) *pb.Application {
	t.Helper()

	if src == nil {
		src = &pb.Application{}
	}

	require.NoError(t, mergo.Merge(src, &pb.Application{
		Project: &pb.Ref_Project{
			Project: "test",
		},

		Name: "test",
	}))

	return src
}

// ValidateUpsertApplicationRequest
func ValidateUpsertApplicationRequest(v *pb.UpsertApplicationRequest) error {
	return validationext.Error(validation.ValidateStruct(v,
		validation.Field(&v.Project, validation.Required),
		validation.Field(&v.Name, validation.Required),
	))
}
