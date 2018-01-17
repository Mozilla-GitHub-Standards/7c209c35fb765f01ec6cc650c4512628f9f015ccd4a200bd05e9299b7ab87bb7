package doorman

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	jwt "gopkg.in/square/go-jose.v2/jwt"
)

func TestMozillaClaimsExtractor(t *testing.T) {
	token, err := jwt.ParseSigned("eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImtpZCI6Ik1rWkRORGN5UmtOR1JURkROamxCTmpaRk9FSkJOMFpCTnpKQlFUTkVNRGhDTUVFd05rRkdPQSJ9.eyJuYW1lIjoiTWF0aGlldSBMZXBsYXRyZSIsImdpdmVuX25hbWUiOiJNYXRoaWV1IiwiZmFtaWx5X25hbWUiOiJMZXBsYXRyZSIsIm5pY2tuYW1lIjoiTWF0aGlldSBMZXBsYXRyZSIsInBpY3R1cmUiOiJodHRwczovL3MuZ3JhdmF0YXIuY29tL2F2YXRhci85NzE5N2YwMTFhM2Q5ZDQ5NGFlODEzNTY2ZjI0Njc5YT9zPTQ4MCZyPXBnJmQ9aHR0cHMlM0ElMkYlMkZjZG4uYXV0aDAuY29tJTJGYXZhdGFycyUyRm1sLnBuZyIsInVwZGF0ZWRfYXQiOiIyMDE3LTEyLTA0VDE1OjUyOjMzLjc2MVoiLCJpc3MiOiJodHRwczovL2F1dGgubW96aWxsYS5hdXRoMC5jb20vIiwic3ViIjoiYWR8TW96aWxsYS1MREFQfG1sZXBsYXRyZSIsImF1ZCI6IlNMb2NmN1NhMWliZDVHTkpNTXFPNTM5ZzdjS3ZXQk9JIiwiZXhwIjoxNTEzMDA3NTcwLCJpYXQiOjE1MTI0MDI3NzAsImFtciI6WyJtZmEiXSwiYWNyIjoiaHR0cDovL3NjaGVtYXMub3BlbmlkLm5ldC9wYXBlL3BvbGljaWVzLzIwMDcvMDYvbXVsdGktZmFjdG9yIiwibm9uY2UiOiJQRkxyLmxtYWhCQWRYaEVSWm0zYVFxc2ZuWjhwcWt0VSIsImF0X2hhc2giOiJTN0Rha1BrZVA0Tnk4SWpTOGxnMHJBIiwiaHR0cHM6Ly9zc28ubW96aWxsYS5jb20vY2xhaW0vZ3JvdXBzIjpbIkludHJhbmV0V2lraSIsIlN0YXRzRGFzaGJvYXJkIiwicGhvbmVib29rX2FjY2VzcyIsImNvcnAtdnBuIiwidnBuX2NvcnAiLCJ2cG5fZGVmYXVsdCIsIkNsb3Vkc2VydmljZXNXaWtpIiwidGVhbV9tb2NvIiwiaXJjY2xvdWQiLCJva3RhX21mYSIsImNsb3Vkc2VydmljZXNfZGV2IiwidnBuX2tpbnRvMV9zdGFnZSIsInZwbl9raW50bzFfcHJvZCIsImVnZW5jaWFfZGUiLCJhY3RpdmVfc2NtX2xldmVsXzEiLCJhbGxfc2NtX2xldmVsXzEiLCJzZXJ2aWNlX3NhZmFyaWJvb2tzIl0sImh0dHBzOi8vc3NvLm1vemlsbGEuY29tL2NsYWltL2VtYWlscyI6WyJtbGVwbGF0cmVAbW96aWxsYS5jb20iLCJtYXRoaWV1QG1vemlsbGEuY29tIiwibWF0aGlldS5sZXBsYXRyZUBtb3ppbGxhLmNvbSJdLCJodHRwczovL3Nzby5tb3ppbGxhLmNvbS9jbGFpbS9kbiI6Im1haWw9bWxlcGxhdHJlQG1vemlsbGEuY29tLG89Y29tLGRjPW1vemlsbGEiLCJodHRwczovL3Nzby5tb3ppbGxhLmNvbS9jbGFpbS9vcmdhbml6YXRpb25Vbml0cyI6Im1haWw9bWxlcGxhdHJlQG1vemlsbGEuY29tLG89Y29tLGRjPW1vemlsbGEiLCJodHRwczovL3Nzby5tb3ppbGxhLmNvbS9jbGFpbS9lbWFpbF9hbGlhc2VzIjpbIm1hdGhpZXVAbW96aWxsYS5jb20iLCJtYXRoaWV1LmxlcGxhdHJlQG1vemlsbGEuY29tIl0sImh0dHBzOi8vc3NvLm1vemlsbGEuY29tL2NsYWltL19IUkRhdGEiOnsicGxhY2Vob2xkZXIiOiJlbXB0eSJ9fQ.MK3Z1Nj15MfbM2TcO4FWVTTYPqAbUhL26pYOFa92mPnEUR2W_oJhwoZ8Vwq7dJcvTZfPq-aZKBnqHoPHHYlQbtaqfflhHmY9iRH0aPlxLQed_WVem4YqMn9xw0az4xHnf0UlzLU58kI97bqUFvvzs0fg_OTdDdO3owVUcaZrG8-xalCqQGQqwTfiH514gxeZ_Ki6610HSVDvpPvmODWPz87IDdgS6WkyM-SyAc3aYukP38aqRo-PUjEdpGbOtV_T_W2x8A3yQDxu0Bcq0WJz-FUEu2BHq1Vn6rmLm7BVYjDD6rYseusp8M0bvTfvXA-9OhJWGAAh6KrN9fnw7r30LQ")
	require.Nil(t, err)

	validator := newJWTGenericValidator("https://auth.mozilla.auth0.com")
	jwks, err := validator.jwks()
	require.Nil(t, err)
	key := &jwks.Keys[0]

	claims, err := mozillaExtractor.Extract(token, key)
	require.Nil(t, err)
	assert.Contains(t, claims.Subject, "|Mozilla-LDAP|")
	assert.Contains(t, claims.Email, "@mozilla.com")
	assert.Contains(t, claims.Groups, "cloudservices_dev", "irccloud")

	// Email provided in `email` field instead of https://sso.../emails list
	token, err = jwt.ParseSigned("eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImtpZCI6Ik1rWkRORGN5UmtOR1JURkROamxCTmpaRk9FSkJOMFpCTnpKQlFUTkVNRGhDTUVFd05rRkdPQSJ9.eyJuYW1lIjoiTWF0aGlldSBMZXBsYXRyZSIsImdpdmVuX25hbWUiOiJNYXRoaWV1IiwiZmFtaWx5X25hbWUiOiJMZXBsYXRyZSIsIm5pY2tuYW1lIjoiTWF0aGlldSBMZXBsYXRyZSIsInBpY3R1cmUiOiJodHRwczovL3MuZ3JhdmF0YXIuY29tL2F2YXRhci85NzE5N2YwMTFhM2Q5ZDQ5NGFlODEzNTY2ZjI0Njc5YT9zPTQ4MCZyPXBnJmQ9aHR0cHMlM0ElMkYlMkZjZG4uYXV0aDAuY29tJTJGYXZhdGFycyUyRm1sLnBuZyIsInVwZGF0ZWRfYXQiOiIyMDE3LTEyLTEzVDIzOjE0OjQ0LjUzOVoiLCJlbWFpbCI6Im1sZXBsYXRyZUBtb3ppbGxhLmNvbSIsImVtYWlsX3ZlcmlmaWVkIjp0cnVlLCJpc3MiOiJodHRwczovL2F1dGgubW96aWxsYS5hdXRoMC5jb20vIiwic3ViIjoiYWR8TW96aWxsYS1MREFQfG1sZXBsYXRyZSIsImF1ZCI6IlNMb2NmN1NhMWliZDVHTkpNTXFPNTM5ZzdjS3ZXQk9JIiwiZXhwIjoxNTEzODExNjg0LCJpYXQiOjE1MTMyMDY4ODQsIm5vbmNlIjoickhOSXF5bGM3SE54MmFhNjktay1SbVA1Y3VqVWNudUkiLCJhdF9oYXNoIjoiZllPZzB6elNHSk1ZWlZTNFRsLXV3dyIsImh0dHBzOi8vc3NvLm1vemlsbGEuY29tL2NsYWltL2dyb3VwcyI6WyJJbnRyYW5ldFdpa2kiLCJTdGF0c0Rhc2hib2FyZCIsInBob25lYm9va19hY2Nlc3MiLCJjb3JwLXZwbiIsInZwbl9jb3JwIiwidnBuX2RlZmF1bHQiLCJDbG91ZHNlcnZpY2VzV2lraSIsInRlYW1fbW9jbyIsImlyY2Nsb3VkIiwib2t0YV9tZmEiLCJjbG91ZHNlcnZpY2VzX2RldiIsInZwbl9raW50bzFfc3RhZ2UiLCJ2cG5fa2ludG8xX3Byb2QiLCJlZ2VuY2lhX2RlIiwiYWN0aXZlX3NjbV9sZXZlbF8xIiwiYWxsX3NjbV9sZXZlbF8xIiwic2VydmljZV9zYWZhcmlib29rcyIsImV2ZXJ5b25lIl0sImh0dHBzOi8vc3NvLm1vemlsbGEuY29tL2NsYWltL1JFQURNRV9GSVJTVCI6IlBsZWFzZSByZWZlciB0byBodHRwczovL2dpdGh1Yi5jb20vbW96aWxsYS1pYW0vcGVyc29uLWFwaSBpbiBvcmRlciB0byBxdWVyeSBNb3ppbGxhIElBTSBDSVMgdXNlciBwcm9maWxlIGRhdGEifQ.EnF3oPHm90ZXnJ4egJqr-4eTaHMw-16beuZlvC66UsIehX7nBooP4VRfMW7KLwOHEnVVGV8jlxgn5p3Dnv1V_W6Yx4PLw7loeKrfhnEKw9onaH3frR_Vo0Y0-MgH4VnCbTwtGHsAfl32j2EoDljXYCqPhYCXD4H25o51lemAoKU3xWamF629FjooyhFTZPVI6JzKkOt39dQjALtXL9EVYRk0ameohHzOT0ZHA57H83FTrPmY_Jy5MWxv1aswcbzcENU1HsFEEkxkRCnGiosxYkStmDo957OQ0IXgNxdNe4VVXzuy5YiNmsjN-IF4tOADLFK5KnLHi4OBOGYiiRiJcQ")
	require.Nil(t, err)
	claims, err = mozillaExtractor.Extract(token, key)
	require.Nil(t, err)
	assert.Contains(t, claims.Email, "@mozilla.com")
}
