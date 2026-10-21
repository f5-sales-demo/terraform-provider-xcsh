---
page_title: "virtual_server.source_port.source_port_preserve_strict"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["virtual server source port source port preserve strict"], "body_bytes": 1245, "body_sha256": "sha256:271afc331060d43e50921ad2160d207155fc5ed612b6ccc7c34c24d918e2ab8e", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:application_profiles:collection", "completeness": "complete", "id": "xcsh-docs:resources:application_profiles:properties:virtual_server:source_port:source_port_preserve_strict", "parent_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:source_port", "path": "documentation/resources/application_profiles/properties/virtual_server/source_port/source_port_preserve_strict/index.md", "product": "distributed-cloud", "provider_name": "application_profiles", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-2110002333222312-2100332330303222-0000003122110232-0110132011311230-1123023003303233-2120121032203202-3211123212022100-3030233302310031", "registry_path": "docs/guides/resources--application_profiles--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["virtual_server", "source_port", "source_port_preserve_strict"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/application_profiles/properties/virtual_server/source_port/source_port_preserve_strict/index.txt", "spec_pin_digest": "sha256:d30f9f9a1cd22b9ba948a7409b4aa6907eda743fb73415c72474bd2784e91af1", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.5", "schema_components": ["application_profilesCreateRequest"], "target_commit": "85874b2a5cfba24c95e5f7d52218df326d4a0582"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# virtual_server.source_port.source_port_preserve_strict

Breadcrumbs:

- [xcsh_application_profiles](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/)
- [virtual_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/)
- [virtual_server.source_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/source_port/)
- virtual_server.source_port.source_port_preserve_strict

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
source_port_preserve_strict = {}
```

This is an empty object or choice marker. It has no direct properties.
