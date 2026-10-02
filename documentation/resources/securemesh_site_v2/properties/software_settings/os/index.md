---
page_title: "software_settings.os"
subcategory: ""
description: "Select the F5XC Operating System Version for the site. By default, latest available OS Version will be used. Refer to release notes to find required released OS versions."
xcsh_docs: {"aliases": ["software settings os"], "body_bytes": 3315, "body_sha256": "sha256:7fa6b9c03cc3476bc0b5bcce61d75c40fc6ca3637880e9bb179cdb7cb7159d4c", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:securemesh_site_v2:properties:software_settings:os:default_os_version"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:software_settings:os", "parent_id": "xcsh-docs:resources:securemesh_site_v2:properties:software_settings", "path": "documentation/resources/securemesh_site_v2/properties/software_settings/os/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0322333322002322-3032123013133021-1301213130312232-1022000103110120-2323021100313310-3330102222002010-2330313032203300-3203333132213302", "registry_path": "docs/guides/resources--securemesh_site_v2--reference--group-017.md", "relationships": [{"anchor": "schema-software_settings--os--operating_system_version", "enforcement": "provider-schema", "group": "software_settings.os:ConflictingObjectAttributes:default_os_version,operating_system_version", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:software_settings:os", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "software_settings.os:ConflictingObjectAttributes:default_os_version,operating_system_version", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:software_settings:os:default_os_version", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["software_settings", "os"], "schema_version": 1, "sections": [{"aliases": ["default os version"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:software_settings:os:default_os_version", "flags": ["optional", "sensitive"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["software_settings", "os", "default_os_version"], "syntax": "attribute", "type": "object"}, {"aliases": ["operating system version"], "anchor": "schema-software_settings--os--operating_system_version", "description": "Exclusive with Specify a OS version to be used e.g. 9.2024.6.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:software_settings:os", "flags": ["optional", "sensitive"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["software_settings", "os", "operating_system_version"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/software_settings/os/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Select the F5XC Operating System Version for the site. By default, latest available OS Version will be used. Refer to release notes to find required released OS versions.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# software_settings.os

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/)
- [software_settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/software_settings/)
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

- [default_os_version](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/software_settings/os/default_os_version/): complete subsection reference.

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

- [software_settings.os.default_os_version](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/software_settings/os/default_os_version/)
- [software_settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/software_settings/)
- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
