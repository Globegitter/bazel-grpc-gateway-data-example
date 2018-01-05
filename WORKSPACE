http_archive(
    name = "io_bazel_rules_go",
    sha256 = "bad35f481c7ca1b3bb3e12d738dfd3f9318073a991fd50d8dc27dffb68b5e23a",
    strip_prefix = "rules_go-396f2424dd7c336e4b07b4f19d476f0749bc73dd",
    url = "https://github.com/bazelbuild/rules_go/archive/396f2424dd7c336e4b07b4f19d476f0749bc73dd.tar.gz",
)

load("@io_bazel_rules_go//go:def.bzl", "go_rules_dependencies", "go_register_toolchains", "go_repository")
load("@io_bazel_rules_go//proto:def.bzl", "proto_register_toolchains")

go_rules_dependencies()
go_register_toolchains()

git_repository(
  name = "org_pubref_rules_protobuf",
  remote = "https://github.com/pubref/rules_protobuf",
  commit = "d5d5a8e30b81599e25f1ff54ea870ab0194b96df" # alternatively, use latest commit on master
)
load("@org_pubref_rules_protobuf//grpc_gateway:rules.bzl", "grpc_gateway_proto_repositories")
grpc_gateway_proto_repositories()

go_repository(
    name = "org_golang_google_grpc",
    commit = "e687fa4e6424368ece6e4fe727cea2c806a0fcb4",
    importpath = "google.golang.org/grpc",
)

go_repository(
    name = "com_github_golang_protobuf",
    commit = "1e59b77b52bf8e4b449a57e6f79f21226d571845",
    importpath = "github.com/golang/protobuf",
)

go_repository(
    name = "com_github_grpc_ecosystem_grpc_gateway",
    importpath = "github.com/grpc-ecosystem/grpc-gateway",
    tag = "v1.3.1",
    build_file_proto_mode="disable",
)

go_repository(
    name = "org_golang_google_genproto",
    commit = "a8101f21cf983e773d0c1133ebc5424792003214",
    importpath = "google.golang.org/genproto",
)
