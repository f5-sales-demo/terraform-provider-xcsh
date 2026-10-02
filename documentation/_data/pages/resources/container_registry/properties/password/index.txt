---
page_title: "password"
subcategory: "Container"
description: "SecretType is used in an object to indicate a sensitive/confidential field."
xcsh_docs: {"aliases": ["password"], "body_bytes": 2036, "body_sha256": "sha256:5f11e8b5ac415f6d3663ffbf3bc287e009d3b3ad126b8d400b735940f6947b61", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:resources:container_registry:properties:password:blindfold_secret_info", "xcsh-docs:resources:container_registry:properties:password:clear_secret_info"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:container_registry:collection", "completeness": "complete", "id": "xcsh-docs:resources:container_registry:properties:password", "parent_id": "xcsh-docs:resources:container_registry:reference", "path": "documentation/resources/container_registry/properties/password/index.md", "product": "distributed-cloud", "provider_name": "container_registry", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2333301020100202-1213333012031101-2112222013010032-1133212101000011-3131130330211132-2113313300202133-3232233203032013-0020230201003030", "registry_path": "docs/guides/resources--container_registry--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "password:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:container_registry:properties:password:blindfold_secret_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "password:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:container_registry:properties:password:clear_secret_info", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["password"], "schema_version": 1, "sections": [{"aliases": ["blindfold secret info"], "anchor": "section", "description": "BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.", "document_id": "xcsh-docs:resources:container_registry:properties:password:blindfold_secret_info", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-password--blindfold_secret_info--location", "enforcement": "provider-schema", "group": "password.blindfold_secret_info:RequiredObjectAttributes:location", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:container_registry:properties:password:blindfold_secret_info", "type": "requires"}], "schema_path": ["password", "blindfold_secret_info"], "syntax": "block", "type": "object"}, {"aliases": ["clear secret info"], "anchor": "section", "description": "ClearSecretInfoType specifies information about the Secret that is not encrypted.", "document_id": "xcsh-docs:resources:container_registry:properties:password:clear_secret_info", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-password--clear_secret_info--url", "enforcement": "provider-schema", "group": "password.clear_secret_info:RequiredObjectAttributes:url", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:container_registry:properties:password:clear_secret_info", "type": "requires"}], "schema_path": ["password", "clear_secret_info"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/container_registry/properties/password/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "SecretType is used in an object to indicate a sensitive/confidential field.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["container_registryCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# password

Breadcrumbs:

- [xcsh_container_registry](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/container_registry/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/container_registry/properties/)
- password

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("blindfold_secret_info",
    "clear_secret_info")}
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
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

Terraform syntax:

```terraform
password {
  # Configure direct properties listed below.
}
```

## Direct properties

- [blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/container_registry/properties/password/blindfold_secret_info/): complete subsection reference.

- [clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/container_registry/properties/password/clear_secret_info/): complete subsection reference.

## Next pages

- [password.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/container_registry/properties/password/blindfold_secret_info/)
- [password.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/container_registry/properties/password/clear_secret_info/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/container_registry/properties/)
- [xcsh_container_registry](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/container_registry/)
