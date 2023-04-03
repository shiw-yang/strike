syntax = "proto3";

package {{.Package}};

option go_package = "{{.GoPackage}}";

service {{.Service}} {
	rpc Create{{.Service}} (Create{{.Service}}Request) returns (Create{{.Service}}Reply);
	rpc Update{{.Service}} (Update{{.Service}}Request) returns (Update{{.Service}}Reply);
	rpc Delete{{.Service}} (Delete{{.Service}}Request) returns (Delete{{.Service}}Reply);
	rpc Get{{.Service}} (Get{{.Service}}Request) returns (Get{{.Service}}Reply);
	rpc List{{.Service}} (List{{.Service}}Request) returns (List{{.Service}}Reply);
}

message Create{{.Service}}Request {}
message Create{{.Service}}Reply {}

message Update{{.Service}}Request {}
message Update{{.Service}}Reply {}

message Delete{{.Service}}Request {}
message Delete{{.Service}}Reply {}

message Get{{.Service}}Request {}
message Get{{.Service}}Reply {}

message List{{.Service}}Request {}
message List{{.Service}}Reply {}