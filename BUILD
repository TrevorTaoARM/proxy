licenses(["notice"])  # Apache 2

load(
    "@envoy//bazel:envoy_build_system.bzl",
    "envoy_cc_binary",
    "envoy_cc_test",
)

exports_files(["linux/bpf_common.h", "linux/bpf.h", "linux/type_mapper.h",
               "proxylib/libcilium.h", "proxylib/types.h"])

envoy_cc_binary(
    name = "cilium-envoy",
    repository = "@envoy",
    visibility = ["//visibility:public"],
    deps = [
        # Cilium filters.
        "//cilium:bpf_metadata_lib",
        "//cilium:network_filter_lib",
        "//cilium:l7policy_lib",
        "//cilium:tls_wrapper_lib",

        "@envoy//source/exe:envoy_main_entry_lib",
    ],
)

envoy_cc_test(
    name = "cilium_integration_test",
    srcs = ["cilium_integration_test.cc"],
    data = [
        "cilium_proxy_test.json",
        "proxylib/libcilium.so",
        "@envoy//test/config/integration/certs",
    ],
    repository = "@envoy",
    deps = [
        "//cilium:bpf_metadata_lib",
        "//cilium:network_filter_lib",
        "//cilium:l7policy_lib",
        "//cilium:tls_wrapper_lib",
        "@envoy//test/integration:http_integration_lib",
        #"@envoy//source/extensions/transport_sockets/tls:context_config_lib",
    ],
)

sh_test(
    name = "envoy_binary_test",
    srcs = ["envoy_binary_test.sh"],
    data = [":cilium-envoy"],
)

sh_binary(
    name = "check_format.py",
    srcs = ["@envoy//tools:check_format.py"],
    deps = [
        ":envoy_build_fixer.py",
        ":header_order.py",
    ],
)

sh_library(
    name = "header_order.py",
    srcs = ["@envoy//tools:header_order.py"],
)

sh_library(
    name = "envoy_build_fixer.py",
    srcs = ["@envoy//tools:envoy_build_fixer.py"],
)
