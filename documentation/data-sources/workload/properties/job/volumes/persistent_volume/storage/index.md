---
page_title: "job.volumes.persistent_volume.storage"
subcategory: "Container"
description: "Persistent storage configuration is used to configure Persistent Volume Claim (PVC)"
xcsh_docs: {"aliases": ["job volumes persistent volume storage"], "body_bytes": 5209, "body_sha256": "sha256:9c86d901288290dbb3afe1d3f03da9d1958245f58b4bb5e130f673a9e190f53b", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:data-sources:workload:properties:job:volumes:persistent_volume:storage:default"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:job:volumes:persistent_volume:storage", "parent_id": "xcsh-docs:data-sources:workload:properties:job:volumes:persistent_volume", "path": "documentation/data-sources/workload/properties/job/volumes/persistent_volume/storage/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-2213202331120023-2231110111323223-1301320301300323-2103321202200000-2113211210103200-1220133202221033-2101020223110213-3203011223010331", "registry_path": "docs/guides/data-sources--workload--reference--group-005.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["job", "volumes", "persistent_volume", "storage"], "schema_version": 1, "sections": [{"aliases": ["job volumes persistent volume storage access mode"], "anchor": "schema-job--volumes--persistent_volume--storage--access_mode", "description": "Persistence storage access mode is used to configure access mode for persistent storage - ACCESS_MODE_READ_WRITE_ONCE: Read Write Once Read Write Once is used to mount persistent storage in read/write mode to exactly 1 host - ACCESS_MODE_READ_WRITE_MANY: Read Write Many Read Write Many is used to mount persistent", "document_id": "xcsh-docs:data-sources:workload:properties:job:volumes:persistent_volume:storage", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["job", "volumes", "persistent_volume", "storage", "access_mode"], "syntax": "attribute", "type": "string"}, {"aliases": ["job volumes persistent volume storage class name"], "anchor": "schema-job--volumes--persistent_volume--storage--class_name", "description": "Exclusive with Use the specified class name.", "document_id": "xcsh-docs:data-sources:workload:properties:job:volumes:persistent_volume:storage", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["job", "volumes", "persistent_volume", "storage", "class_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["job volumes persistent volume storage default"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:workload:properties:job:volumes:persistent_volume:storage:default", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["job", "volumes", "persistent_volume", "storage", "default"], "syntax": "attribute", "type": "object"}, {"aliases": ["job volumes persistent volume storage storage size"], "anchor": "schema-job--volumes--persistent_volume--storage--storage_size", "description": "Size in GiB of the persistent storage.", "document_id": "xcsh-docs:data-sources:workload:properties:job:volumes:persistent_volume:storage", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["job", "volumes", "persistent_volume", "storage", "storage_size"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/job/volumes/persistent_volume/storage/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Persistent storage configuration is used to configure Persistent Volume Claim (PVC)", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# job.volumes.persistent_volume.storage

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/)
- [job](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/job/)
- [job.volumes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/job/volumes/)
- [job.volumes.persistent_volume](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/job/volumes/persistent_volume/)
- job.volumes.persistent_volume.storage

<a id="section"></a>

Type: `"single"`. Computed.

Persistent storage configuration is used to configure Persistent Volume Claim (PVC).

Upstream description:

Persistent storage configuration is used to configure Persistent Volume Claim (PVC)

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-class_name_choice": "[\"class_name\",\"default\"]"
}
```

## Direct properties

<a id="schema-job--volumes--persistent_volume--storage--access_mode"></a>

### access_mode property

Type: `"string"`. Computed.

\[Enum:
ACCESS\_MODE\_READ\_WRITE\_ONCE|ACCESS\_MODE\_READ\_WRITE\_MANY|ACCESS\_MODE\_READ\_ONLY\_MANY\]
Persistence storage access mode is used to configure access mode for persistent storage -
ACCESS\_MODE\_READ\_WRITE\_ONCE: Read Write Once Read Write Once is used to mount persistent storage
in read/write mode to exactly 1 host - ACCESS\_MODE\_READ\_WRITE\_MANY: Read Write Many Read Write
Many is used.. Possible values are \`ACCESS\_MODE\_READ\_WRITE\_ONCE\`,
\`ACCESS\_MODE\_READ\_WRITE\_MANY\`, \`ACCESS\_MODE\_READ\_ONLY\_MANY\`. Defaults to
\`ACCESS\_MODE\_READ\_WRITE\_ONCE\`.

Upstream description:

Persistence storage access mode is used to configure access mode for persistent storage

&#8203;- ACCESS\_MODE\_READ\_WRITE\_ONCE: Read Write Once

Read Write Once is used to mount persistent storage in read/write mode to exactly 1 host &#8203;-
ACCESS\_MODE\_READ\_WRITE\_MANY: Read Write Many

Read Write Many is used to mount persistent storage in read/write mode to many hosts &#8203;-
ACCESS\_MODE\_READ\_ONLY\_MANY: Read Only Many

Read Only Many is used to mount persistent storage in read-only mode to many hosts.

Receipt-pinned upstream constraints:

```json
{
  "default": "ACCESS_MODE_READ_WRITE_ONCE",
  "enum": [
    "ACCESS_MODE_READ_WRITE_ONCE",
    "ACCESS_MODE_READ_WRITE_MANY",
    "ACCESS_MODE_READ_ONLY_MANY"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-job--volumes--persistent_volume--storage--class_name"></a>

### class_name property

Type: `"string"`. Computed.

Exclusive with \[default\] Use the specified class name.

Upstream description:

Exclusive with \[default\] Use the specified class name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [default](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/job/volumes/persistent_volume/storage/default/): complete subsection reference.

<a id="schema-job--volumes--persistent_volume--storage--storage_size"></a>

### storage_size property

Type: `"number"`. Computed.

Size (in GiB). Size in GiB of the persistent storage.

Upstream description:

Size in GiB of the persistent storage.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.double.gte": "0.004",
    "ves.io.schema.rules.double.lte": "256",
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.double.gte": "0.004",
    "ves.io.schema.rules.double.lte": "256",
    "ves.io.schema.rules.message.required": "true"
  }
}
```

## Next pages

- [job.volumes.persistent_volume.storage.default](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/job/volumes/persistent_volume/storage/default/)
- [job.volumes.persistent_volume](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/job/volumes/persistent_volume/)
- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
