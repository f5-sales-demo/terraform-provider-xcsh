---
page_title: "ddos_profile.enable_ddos_mitigation"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["ddos profile enable ddos mitigation"], "body_bytes": 996, "body_sha256": "sha256:7001c31efb50469f16b284351f59196cd31f6ddb3271786094360936433ee661", "capabilities": ["dns"], "category": "dns", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:dns_proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_proxy:properties:ddos_profile:enable_ddos_mitigation", "parent_id": "xcsh-docs:resources:dns_proxy:properties:ddos_profile", "path": "documentation/resources/dns_proxy/properties/ddos_profile/enable_ddos_mitigation/index.md", "product": "distributed-cloud", "provider_name": "dns_proxy", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "resources", "registry_anchor": "canonical-0113322033110122-3010012031011023-0120030310322113-1011300230111202-3001303212100313-3020111302222203-2333101200222321-0223001130313013", "registry_path": "docs/guides/resources--dns_proxy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["ddos_profile", "enable_ddos_mitigation"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_proxy/properties/ddos_profile/enable_ddos_mitigation/index.txt", "spec_pin_digest": "sha256:66dcd7f1752e908b5a58a7638e63662f0daa121544d954167e4fab269122ff4a", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.1", "schema_components": ["dns_proxyCreateRequest"], "target_commit": "bafefa99cbad3ad2e3c2c4fd9c0611a0cbcc36b3"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ddos_profile.enable_ddos_mitigation

Breadcrumbs:

- [xcsh_dns_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/)
- [ddos_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/ddos_profile/)
- ddos_profile.enable_ddos_mitigation

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
enable_ddos_mitigation = {}
```

This is an empty object or choice marker. It has no direct properties.
