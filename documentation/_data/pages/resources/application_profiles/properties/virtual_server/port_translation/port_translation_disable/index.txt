---
page_title: "virtual_server.port_translation.port_translation_disable"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["virtual server port translation port translation disable"], "body_bytes": 1256, "body_sha256": "sha256:ccd2d44af82b759230785a73c9619d4bf51b6c2b97eca9dfb1f09b2320695875", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:application_profiles:collection", "completeness": "complete", "id": "xcsh-docs:resources:application_profiles:properties:virtual_server:port_translation:port_translation_disable", "parent_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:port_translation", "path": "documentation/resources/application_profiles/properties/virtual_server/port_translation/port_translation_disable/index.md", "product": "distributed-cloud", "provider_name": "application_profiles", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "resources", "registry_anchor": "canonical-2322122032001202-0320333323113103-1320111022100231-1122220021233203-0132331132331011-0222313003233312-0311310233221122-0300233002300010", "registry_path": "docs/guides/resources--application_profiles--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["virtual_server", "port_translation", "port_translation_disable"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/application_profiles/properties/virtual_server/port_translation/port_translation_disable/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["application_profilesCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# virtual_server.port_translation.port_translation_disable

Breadcrumbs:

- [xcsh_application_profiles](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/)
- [virtual_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/)
- [virtual_server.port_translation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/port_translation/)
- virtual_server.port_translation.port_translation_disable

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
port_translation_disable = {}
```

This is an empty object or choice marker. It has no direct properties.
