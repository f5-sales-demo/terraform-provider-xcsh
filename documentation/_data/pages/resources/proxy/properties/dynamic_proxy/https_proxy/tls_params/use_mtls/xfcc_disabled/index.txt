---
page_title: "dynamic_proxy.https_proxy.tls_params.use_mtls.xfcc_disabled"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["dynamic proxy https proxy tls params use mtls xfcc disabled"], "body_bytes": 1506, "body_sha256": "sha256:6fe53b9b3d8690f27c2813b2491a3e080b63fed5a0fe70b58079d3400bd505c5", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy:tls_params:use_mtls:xfcc_disabled", "parent_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy:tls_params:use_mtls", "path": "documentation/resources/proxy/properties/dynamic_proxy/https_proxy/tls_params/use_mtls/xfcc_disabled/index.md", "product": "distributed-cloud", "provider_name": "proxy", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-3301301013003012-1010130000202121-0230031303003100-2111102033131312-0200320003300232-1300031313132231-1320223300221001-0012021223131322", "registry_path": "docs/guides/resources--proxy--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["dynamic_proxy", "https_proxy", "tls_params", "use_mtls", "xfcc_disabled"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/proxy/properties/dynamic_proxy/https_proxy/tls_params/use_mtls/xfcc_disabled/index.txt", "spec_pin_digest": "sha256:ea0d80973b5efe05eb3575be6beda9c96b86701d264dc1aee3972e6a23ebb3fb", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.3", "schema_components": ["proxyCreateRequest"], "target_commit": "6ee1e85b2e917d58078f38a270b790077b3fba12"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# dynamic_proxy.https_proxy.tls_params.use_mtls.xfcc_disabled

Breadcrumbs:

- [xcsh_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/)
- [dynamic_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/dynamic_proxy/)
- [dynamic_proxy.https_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/dynamic_proxy/https_proxy/)
- [dynamic_proxy.https_proxy.tls_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/dynamic_proxy/https_proxy/tls_params/)
- [dynamic_proxy.https_proxy.tls_params.use_mtls](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/dynamic_proxy/https_proxy/tls_params/use_mtls/)
- dynamic_proxy.https_proxy.tls_params.use_mtls.xfcc_disabled

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

Additional upstream details:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
xfcc_disabled = {}
```

This is an empty object or choice marker. It has no direct properties.
