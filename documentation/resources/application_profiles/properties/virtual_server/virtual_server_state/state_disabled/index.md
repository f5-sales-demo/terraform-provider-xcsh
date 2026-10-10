---
page_title: "virtual_server.virtual_server_state.state_disabled"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["virtual server virtual server state state disabled"], "body_bytes": 1242, "body_sha256": "sha256:bdd5129268b54a50f84ab10a44e95aff2aa4d5047a0442e8d0846cff5616547a", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:application_profiles:collection", "completeness": "complete", "id": "xcsh-docs:resources:application_profiles:properties:virtual_server:virtual_server_state:state_disabled", "parent_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:virtual_server_state", "path": "documentation/resources/application_profiles/properties/virtual_server/virtual_server_state/state_disabled/index.md", "product": "distributed-cloud", "provider_name": "application_profiles", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-0121331000132220-1313223010021111-1300122021210133-2121010231300102-3330300311303212-0031200013322333-1123123322332330-2010212310333330", "registry_path": "docs/guides/resources--application_profiles--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["virtual_server", "virtual_server_state", "state_disabled"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/application_profiles/properties/virtual_server/virtual_server_state/state_disabled/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["application_profilesCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# virtual_server.virtual_server_state.state_disabled

Breadcrumbs:

- [xcsh_application_profiles](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/)
- [virtual_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/)
- [virtual_server.virtual_server_state](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/virtual_server_state/)
- virtual_server.virtual_server_state.state_disabled

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
state_disabled = {}
```

This is an empty object or choice marker. It has no direct properties.
