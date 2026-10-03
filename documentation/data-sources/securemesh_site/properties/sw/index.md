---
page_title: "sw"
subcategory: ""
description: "Select the F5XC Software Version for the site. By default, latest available F5XC Software Version will be used. Refer to release notes to find required released SW versions."
xcsh_docs: {"aliases": ["sw"], "body_bytes": 2638, "body_sha256": "sha256:bbd0c01296898920a761e99d73132e02e381cf51ea855130e9395d140be9ddf1", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:securemesh_site:properties:sw:default_sw_version"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:securemesh_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site:properties:sw", "parent_id": "xcsh-docs:data-sources:securemesh_site:reference", "path": "documentation/data-sources/securemesh_site/properties/sw/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-1101033231030033-1231112223010232-0121200330001132-1331111222212213-3023302310111332-0331233101000132-0230010020103302-3320133222123131", "registry_path": "docs/guides/data-sources--securemesh_site--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["sw"], "schema_version": 1, "sections": [{"aliases": ["sw default sw version"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:securemesh_site:properties:sw:default_sw_version", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["sw", "default_sw_version"], "syntax": "attribute", "type": "object"}, {"aliases": ["sw volterra software version"], "anchor": "schema-sw--volterra_software_version", "description": "Exclusive with Specify a F5XC Software Version to be used e.g. Crt-20210329-1002.", "document_id": "xcsh-docs:data-sources:securemesh_site:properties:sw", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["sw", "volterra_software_version"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site/properties/sw/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Select the F5XC Software Version for the site. By default, latest available F5XC Software Version will be used. Refer to release notes to find required released SW versions.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["securemesh_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# sw

Breadcrumbs:

- [xcsh_securemesh_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/)
- sw

<a id="section"></a>

Type: `"single"`. Computed.

Select the F5XC Software Version for the site. By default, latest available F5XC Software Version
will be used. Refer to release notes to find required released SW versions.

Upstream description:

Select the F5XC Software Version for the site. By default, latest available F5XC Software Version
will be used. Refer to release notes to find required released SW versions.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-volterra_sw_version_choice": "[\"default_sw_version\",\"volterra_software_version\"]"
}
```

## Direct properties

- [default_sw_version](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/sw/default_sw_version/): complete subsection reference.

<a id="schema-sw--volterra_software_version"></a>

### volterra_software_version property

Type: `"string"`. Computed.

Exclusive with \[default\_sw\_version\] Specify a F5XC Software Version to be used e.g.
Crt-20210329-1002.

Upstream description:

Exclusive with \[default\_sw\_version\] Specify a F5XC Software Version to be used e.g.
Crt-20210329-1002.

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

- [sw.default_sw_version](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/sw/default_sw_version/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/)
- [xcsh_securemesh_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/)
