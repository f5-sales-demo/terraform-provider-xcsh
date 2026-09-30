---
page_title: "software_settings.os"
subcategory: ""
description: "software_settings.os for xcsh_securemesh_site_v2."
xcsh_docs: {"aliases": [], "body_bytes": 2861, "body_sha256": "sha256:485ce9f6be2e620d1f60d37a0964211fc84deef3b4ffbc7d534cd737afb76e2d", "canonical_id": "xcsh-docs:resources:securemesh_site_v2:properties:software_settings:os", "child_ids": ["xcsh-docs:resources:securemesh_site_v2:properties:software_settings:os:default_os_version"], "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:software_settings:os", "parent_id": "xcsh-docs:resources:securemesh_site_v2:properties:software_settings", "path": "docs/guides/resources--securemesh_site_v2--properties--software_settings--os.md", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["software_settings", "os"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/software_settings/os/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "software_settings.os for xcsh_securemesh_site_v2.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# software_settings.os

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md)
- [Property reference](resources--securemesh_site_v2--reference.md)
- [software_settings](resources--securemesh_site_v2--properties--software_settings.md)
- software_settings.os

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Select the F5XC Operating System Version for the site. By default, latest available OS Version will
be used. Refer to release notes to find required released OS versions.

Upstream description:

Select the F5XC Operating System Version for the site. By default, latest available OS Version will
be used. Refer to release notes to find required released OS versions.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("default_os_version",
    "operating_system_version")}
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
  "x-ves-oneof-field-operating_system_version_choice": "[\"default_os_version\",\"operating_system_version\"]"
}
```

Terraform syntax:

```terraform
os {
  # Configure direct properties listed below.
}
```

## Direct properties

- [default_os_version](resources--securemesh_site_v2--properties--software_settings--os--default_os_version.md): complete subsection reference.

<a id="schema-software_settings--os--operating_system_version"></a>

### operating_system_version property

Type: `"string"`. Optional, Sensitive.

Exclusive with \[default\_os\_version\] Specify a OS version to be used e.g. 9.2024.6.

Upstream description:

Exclusive with \[default\_os\_version\] Specify a OS version to be used e.g. 9.2024.6.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(20),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 20,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 20,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "20"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "20"
  }
}
```

## Next pages

- [software_settings.os.default_os_version](resources--securemesh_site_v2--properties--software_settings--os--default_os_version.md)
- [software_settings](resources--securemesh_site_v2--properties--software_settings.md)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md)
