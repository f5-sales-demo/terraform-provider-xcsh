---
page_title: "virtual_server.nat64.nat64_disable"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["virtual server nat64 nat64 disable"], "body_bytes": 1203, "body_sha256": "sha256:28a08782390df774161dcbd9b67a482a82d3864c7477187d8abd71cd4a6ab885", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:application_profiles:collection", "completeness": "complete", "id": "xcsh-docs:resources:application_profiles:properties:virtual_server:nat64:nat64_disable", "parent_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:nat64", "path": "documentation/resources/application_profiles/properties/virtual_server/nat64/nat64_disable/index.md", "product": "distributed-cloud", "provider_name": "application_profiles", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "resources", "registry_anchor": "canonical-1101131332320313-1030303021303121-3201103010330112-0200130321003203-3233323021222323-0021303201111130-0000313032310333-0112133213301313", "registry_path": "docs/guides/resources--application_profiles--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["virtual_server", "nat64", "nat64_disable"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/application_profiles/properties/virtual_server/nat64/nat64_disable/index.txt", "spec_pin_digest": "sha256:d5e8f282045e20f8be00960c9e0ed5287ab3532b8da7d4b559f1cea17fd51bca", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.9", "schema_components": ["application_profilesCreateRequest"], "target_commit": "e04c5d45f7a4e103e2dc84bf52b767182b4de3bc"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# virtual_server.nat64.nat64_disable

Breadcrumbs:

- [xcsh_application_profiles](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/)
- [virtual_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/)
- [virtual_server.nat64](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/nat64/)
- virtual_server.nat64.nat64_disable

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for nat64 disable.

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
nat64_disable = {}
```

This is an empty object or choice marker. It has no direct properties.
