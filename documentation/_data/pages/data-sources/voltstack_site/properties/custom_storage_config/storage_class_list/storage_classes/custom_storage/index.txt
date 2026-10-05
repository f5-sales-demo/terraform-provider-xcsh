---
page_title: "custom_storage_config.storage_class_list.storage_classes.custom_storage"
subcategory: ""
description: "Custom Storage Class allows to insert Kubernetes storageclass definition which will be applied into given site."
xcsh_docs: {"aliases": ["custom storage config storage class list storage classes custom storage"], "body_bytes": 2874, "body_sha256": "sha256:599aeb13f08b249e1c49275b855d3064cf10ccc81a4589d4b0807a31f9de9d68", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_class_list:storage_classes:custom_storage", "parent_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_class_list:storage_classes", "path": "documentation/data-sources/voltstack_site/properties/custom_storage_config/storage_class_list/storage_classes/custom_storage/index.md", "product": "distributed-cloud", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-1322100112301000-3320130211010212-1211322301011211-1201121020333033-1132123112232220-3201231101201010-2102133333021022-1132203203221322", "registry_path": "docs/guides/data-sources--voltstack_site--reference--group-006.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["custom_storage_config", "storage_class_list", "storage_classes", "custom_storage"], "schema_version": 1, "sections": [{"aliases": ["custom storage config storage class list storage classes custom storage yaml"], "anchor": "schema-custom_storage_config--storage_class_list--storage_classes--custom_storage--yaml", "description": "K8s YAML for StorageClass.", "document_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_class_list:storage_classes:custom_storage", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_storage_config", "storage_class_list", "storage_classes", "custom_storage", "yaml"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/voltstack_site/properties/custom_storage_config/storage_class_list/storage_classes/custom_storage/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Custom Storage Class allows to insert Kubernetes storageclass definition which will be applied into given site.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_storage_config.storage_class_list.storage_classes.custom_storage

Breadcrumbs:

- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/)
- [custom_storage_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_storage_config/)
- [custom_storage_config.storage_class_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_storage_config/storage_class_list/)
- [custom_storage_config.storage_class_list.storage_classes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_storage_config/storage_class_list/storage_classes/)
- custom_storage_config.storage_class_list.storage_classes.custom_storage

<a id="section"></a>

Type: `"single"`. Computed.

Custom Storage Class allows to insert Kubernetes storageclass definition which will be applied into
given site.

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

## Direct properties

<a id="schema-custom_storage_config--storage_class_list--storage_classes--custom_storage--yaml"></a>

### yaml property

Type: `"string"`. Computed.

Storage Class YAML. K8s YAML for StorageClass.

Upstream description:

K8s YAML for StorageClass.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 4096,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "formatDescription": "Valid parseable YAML",
    "maxLength": 4096,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "validation": {
      "customRule": "Must be valid YAML"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "4096",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "4096",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

## Next pages

- [custom_storage_config.storage_class_list.storage_classes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_storage_config/storage_class_list/storage_classes/)
- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/)
