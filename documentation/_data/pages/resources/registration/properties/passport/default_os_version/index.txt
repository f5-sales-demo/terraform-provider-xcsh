---
page_title: "passport.default_os_version"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["passport default os version"], "body_bytes": 980, "body_sha256": "sha256:5059eed106b8b6e8ba3c7d61ef81a79d1cda4e2d284394a11da7f0644d842ec6", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:registration:collection", "completeness": "complete", "id": "xcsh-docs:resources:registration:properties:passport:default_os_version", "parent_id": "xcsh-docs:resources:registration:properties:passport", "path": "documentation/resources/registration/properties/passport/default_os_version/index.md", "product": "distributed-cloud", "provider_name": "registration", "provider_schema_digest": "sha256:e62318985676da5dc8f5d2d35ccc26c58c5f11f463e5627995ce95f8658f6b43", "provider_type": "resources", "registry_anchor": "canonical-0121101322021220-2201332112213101-2230331311023230-1311323103322112-0302220111230211-0313313133223230-2003203200110302-0331003203032000", "registry_path": "docs/guides/resources--registration--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["passport", "default_os_version"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/registration/properties/passport/default_os_version/index.txt", "spec_pin_digest": "sha256:b750ce4b9c0ad9c90dad992a893773bacfb3aeff7598bc37194f3037f353c949", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.5.2", "schema_components": ["registrationCreateRequest"], "target_commit": "18cc909b5ddc97ef9abd48764589d753290a31a4"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# passport.default_os_version

Breadcrumbs:

- [xcsh_registration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/properties/)
- [passport](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/properties/passport/)
- passport.default_os_version

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
default_os_version = {}
```

This is an empty object or choice marker. It has no direct properties.
