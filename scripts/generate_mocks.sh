mockery --recursive=true --case=underscore --output=pkg/internal/mocks --name=Repository --structname=MockUserRepo --filename=mock_user_repo.go --dir=pkg/user
mockery --recursive=true --case=underscore --output=pkg/internal/mocks --name=Servicer --structname=MockUserService --filename=mock_user_service.go --dir=pkg/user
mockery --recursive=true --case=underscore --output=pkg/internal/mocks --name=Handler --structname=MockHTTPHandler --filename=mock_http_handler.go --dir=/usr/local/go/src/net/http
