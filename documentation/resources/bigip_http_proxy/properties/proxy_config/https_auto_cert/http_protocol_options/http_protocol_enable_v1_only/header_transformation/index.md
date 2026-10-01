---
page_title: "proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation"
subcategory: ""
description: "proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation for xcsh_bigip_http_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 4771, "body_sha256": "sha256:7def075e68a48965007a5c31811fc59b030c04d6cd09c4b19d7359600b393f48", "child_ids": ["xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https_auto_cert:http_protocol_options:http_protocol_enable_v1_only:header_transformation:default_header_transformation", "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https_auto_cert:http_protocol_options:http_protocol_enable_v1_only:header_transformation:preserve_case_header_transformation", "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https_auto_cert:http_protocol_options:http_protocol_enable_v1_only:header_transformation:proper_case_header_transformation"], "collection_id": "xcsh-docs:resources:bigip_http_proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https_auto_cert:http_protocol_options:http_protocol_enable_v1_only:header_transformation", "parent_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https_auto_cert:http_protocol_options:http_protocol_enable_v1_only", "path": "documentation/resources/bigip_http_proxy/properties/proxy_config/https_auto_cert/http_protocol_options/http_protocol_enable_v1_only/header_transformation/index.md", "provider_name": "bigip_http_proxy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["proxy_config", "https_auto_cert", "http_protocol_options", "http_protocol_enable_v1_only", "header_transformation"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bigip_http_proxy/properties/proxy_config/https_auto_cert/http_protocol_options/http_protocol_enable_v1_only/header_transformation/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation for xcsh_bigip_http_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bigip_http_proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation

Breadcrumbs:

- [xcsh_bigip_http_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/)
- [proxy_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_config/)
- [proxy_config.https_auto_cert](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_config/https_auto_cert/)
- [proxy_config.https_auto_cert.http_protocol_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_config/https_auto_cert/http_protocol_options/)
- [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_config/https_auto_cert/http_protocol_options/http_protocol_enable_v1_only/)
- proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Header Transformation OPTIONS for HTTP/1.1 request/response headers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("default_header_transformation",
    "preserve_case_header_transformation"),
  validators.ConflictingObjectAttributes("default_header_transformation",
    "proper_case_header_transformation"),
  validators.ConflictingObjectAttributes("preserve_case_header_transformation",
    "proper_case_header_transformation")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-header_transformation_choice": "[\"default_header_transformation\",\"preserve_case_header_transformation\",\"proper_case_header_transformation\"]"
}
```

Terraform syntax:

```terraform
header_transformation {
  # Configure direct properties listed below.
}
```

## Direct properties

- [default_header_transformation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_config/https_auto_cert/http_protocol_options/http_protocol_enable_v1_only/header_transformation/default_header_transformation/): complete subsection reference.

- [preserve_case_header_transformation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_config/https_auto_cert/http_protocol_options/http_protocol_enable_v1_only/header_transformation/preserve_case_header_transformation/): complete subsection reference.

- [proper_case_header_transformation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_config/https_auto_cert/http_protocol_options/http_protocol_enable_v1_only/header_transformation/proper_case_header_transformation/): complete subsection reference.

## Next pages

- [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_config/https_auto_cert/http_protocol_options/http_protocol_enable_v1_only/header_transformation/default_header_transformation/)
- [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_config/https_auto_cert/http_protocol_options/http_protocol_enable_v1_only/header_transformation/preserve_case_header_transformation/)
- [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_config/https_auto_cert/http_protocol_options/http_protocol_enable_v1_only/header_transformation/proper_case_header_transformation/)
- [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_config/https_auto_cert/http_protocol_options/http_protocol_enable_v1_only/)
- [xcsh_bigip_http_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/)
