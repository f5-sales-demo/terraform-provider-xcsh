---
page_title: "private_key"
subcategory: "Security"
description: "private_key for xcsh_certificate."
xcsh_docs: {"aliases": [], "body_bytes": 1585, "body_sha256": "sha256:720dd2d070874d1abaea466caaa5649befcc1690dc146759980ce54b17eae5b0", "canonical_id": "xcsh-docs:resources:certificate:properties:private_key", "child_ids": ["xcsh-docs:resources:certificate:properties:private_key:blindfold_secret_info", "xcsh-docs:resources:certificate:properties:private_key:clear_secret_info"], "collection_id": "xcsh-docs:resources:certificate:collection", "completeness": "complete", "id": "xcsh-docs:resources:certificate:properties:private_key", "parent_id": "xcsh-docs:resources:certificate:reference", "path": "docs/guides/resources--certificate--properties--private_key.md", "provider_name": "certificate", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["private_key"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/certificate/properties/private_key/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "private_key for xcsh_certificate.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["certificateCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# private_key

Breadcrumbs:

- [xcsh_certificate](../resources/certificate.md)
- [Property reference](resources--certificate--reference.md)
- private_key

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
private_key {
  # Configure direct properties listed below.
}
```

## Direct properties

- [blindfold_secret_info](resources--certificate--properties--private_key--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](resources--certificate--properties--private_key--clear_secret_info.md): complete subsection reference.

## Next pages

- [private_key.blindfold_secret_info](resources--certificate--properties--private_key--blindfold_secret_info.md)
- [private_key.clear_secret_info](resources--certificate--properties--private_key--clear_secret_info.md)
- [Property reference](resources--certificate--reference.md)
- [xcsh_certificate](../resources/certificate.md)
