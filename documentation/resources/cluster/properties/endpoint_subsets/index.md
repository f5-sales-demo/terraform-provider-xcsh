---
page_title: "endpoint_subsets"
subcategory: ""
description: "Cluster may be configured to divide its endpoints into subsets based on metadata attached to the endpoints. Routes may then specify the metadata that a endpoint must match in order to be selected by the load balancer. Endpoint_subsets is list of subsets for this cluster. Each entry in this list has definition for a"
xcsh_docs: {"aliases": ["endpoint subsets"], "body_bytes": 4818, "body_sha256": "sha256:bf709e66660eff10c1f68ddd9add3c05e32f11272c8aa33dc8d43feccc8417f4", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:cluster:collection", "completeness": "complete", "id": "xcsh-docs:resources:cluster:properties:endpoint_subsets", "parent_id": "xcsh-docs:resources:cluster:reference", "path": "documentation/resources/cluster/properties/endpoint_subsets/index.md", "product": "distributed-cloud", "provider_name": "cluster", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-0130002222020001-0232133310310030-3303113101232230-1302101000000023-2102221011311020-3031020222211032-2220002000132123-3001122011123303", "registry_path": "docs/guides/resources--cluster--reference--group-001.md", "relationships": [{"anchor": "schema-endpoint_subsets--keys", "enforcement": "provider-schema", "group": "endpoint_subsets:RequiredListObjectAttributes:keys", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:cluster:properties:endpoint_subsets", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["endpoint_subsets"], "schema_version": 1, "sections": [{"aliases": ["endpoint subsets keys"], "anchor": "schema-endpoint_subsets--keys", "description": "List of keys that define a cluster subset class.", "document_id": "xcsh-docs:resources:cluster:properties:endpoint_subsets", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["endpoint_subsets", "keys"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cluster/properties/endpoint_subsets/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Cluster may be configured to divide its endpoints into subsets based on metadata attached to the endpoints. Routes may then specify the metadata that a endpoint must match in order to be selected by the load balancer. Endpoint_subsets is list of subsets for this cluster. Each entry in this list has definition for a", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["clusterCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# endpoint_subsets

Breadcrumbs:

- [xcsh_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/)
- endpoint_subsets

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Configure endpoint groups based on metadata labels for traffic routing. Supports weighted
distribution and session affinity across labeled endpoints.

Additional upstream details:

Cluster may be configured to divide its endpoints into subsets based on metadata attached to the
endpoints. Routes may then specify the metadata that a endpoint must match in order to be selected
by the load balancer. Endpoint\_subsets is list of subsets for this cluster. Each entry in this list
has definition for a subset (which is collection of keys)

During routing, the route’s metadata match configuration is used to find a specific subset. If there
is a subset with the exact keys and values specified by the route, the subset is used for load
balancing. Otherwise, the fallback policy is used. The cluster’s subset configuration must,
therefore, contain a definition that has the same keys as a given route in order for subset load
balancing to occur. Example:

RouteConfig

routes: &#8203;- match: &#8203;- headers: \[\] path: path: /1.log query\_params: \[\]
routeDestination: destinations: &#8203;- cluster: &#8203;- kind: cluster.object uid:
00000000-0000-4000-8000-0b50b89d07a2 endpointSubsets: site: india

EndpointConfig

metadata: labels: deployment: debug site: india name: end-1 uid: end-1

ClusterConfig

gcSpec: defaultSubset: stage: production fallbackPolicy: DEFAULT\_SUBSET endpointSubsets: &#8203;-
keys: &#8203;- site &#8203;- keys: &#8203;- stage &#8203;- app

Assume the below endpoints are defined and associated with the cluster. Endpoint Labels --------
&#8203;------

ep1 stage: production, site: india ep2 stage: deployment, site: us ep3 stage: production, app: hr
ep4 site: india

The following table describes some routes and the result of their application to the cluster. The
subset definition for cluster is assumed to be same as given above in the ClusterConfig section

RouteMatch Criteria Subset Reason ------------------- ------ ------

site: india ep1, ep4 Subset of endpoints selected site: us ep2 Subset of endpoints selected app: hr
ep1, ep3 Fallback: No subset selector for "app" alone stage: production, app: hr ep3 Subset of
endpoints selected other: x ep1, ep3 Fallback: No subset selector for “other” (none) ep1, ep3
Fallback: No subset requested.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("keys")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32"
  }
}
```

Terraform syntax:

```terraform
endpoint_subsets {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-endpoint_subsets--keys"></a>

### keys property

Type: `["list", "string"]`. Optional.

List of keys that define a cluster subset class.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```
