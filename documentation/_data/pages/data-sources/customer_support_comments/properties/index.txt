---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_customer_support_comments."
xcsh_docs: {"aliases": [], "body_bytes": 4750, "body_sha256": "sha256:c46772b20683de3243892640d7d5b6941b7c55b0d3dc42dd250797c1e526fb0d", "child_ids": ["xcsh-docs:data-sources:customer_support_comments:properties:comments"], "collection_id": "xcsh-docs:data-sources:customer_support_comments:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:customer_support_comments:reference", "parent_id": "xcsh-docs:data-sources:customer_support_comments:fundamentals", "path": "documentation/data-sources/customer_support_comments/properties/index.md", "provider_name": "customer_support_comments", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/customer_support_comments/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_customer_support_comments.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_customer_support_comments](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/customer_support_comments/)
- Property reference

## Direct properties

<a id="schema-all_comments_returned"></a>

### all_comments_returned property

Type: `"bool"`. Computed.

Indicates if all comments for the issue have been returned.

- [comments](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/customer_support_comments/properties/comments/): complete subsection reference.

<a id="schema-created_until_timestamp"></a>

### created_until_timestamp property

Type: `"string"`. Optional.

Filter to retrieve comments created up to the specified timestamp.

<a id="schema-name"></a>

### name property

Type: `"string"`. Required.

Name The name (issue ID) of the customer support ticket object.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `all_comments_returned` | [all_comments_returned](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/customer_support_comments/properties/#schema-all_comments_returned) |
| `comments` | [comments](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/customer_support_comments/properties/comments/#section) |
| `comments.attachment_ids` | [comments.attachment_ids](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/customer_support_comments/properties/comments/#schema-comments--attachment_ids) |
| `comments.attachments_info` | [comments.attachments_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/customer_support_comments/properties/comments/attachments_info/#section) |
| `comments.attachments_info.attachment` | [comments.attachments_info.attachment](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/customer_support_comments/properties/comments/attachments_info/#schema-comments--attachments_info--attachment) |
| `comments.attachments_info.content_type` | [comments.attachments_info.content_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/customer_support_comments/properties/comments/attachments_info/#schema-comments--attachments_info--content_type) |
| `comments.attachments_info.filename` | [comments.attachments_info.filename](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/customer_support_comments/properties/comments/attachments_info/#schema-comments--attachments_info--filename) |
| `comments.attachments_info.tp_id` | [comments.attachments_info.tp_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/customer_support_comments/properties/comments/attachments_info/#schema-comments--attachments_info--tp_id) |
| `comments.author_email` | [comments.author_email](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/customer_support_comments/properties/comments/#schema-comments--author_email) |
| `comments.author_name` | [comments.author_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/customer_support_comments/properties/comments/#schema-comments--author_name) |
| `comments.comment_id` | [comments.comment_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/customer_support_comments/properties/comments/#schema-comments--comment_id) |
| `comments.created_at` | [comments.created_at](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/customer_support_comments/properties/comments/#schema-comments--created_at) |
| `comments.html` | [comments.html](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/customer_support_comments/properties/comments/#schema-comments--html) |
| `comments.plain_text` | [comments.plain_text](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/customer_support_comments/properties/comments/#schema-comments--plain_text) |
| `created_until_timestamp` | [created_until_timestamp](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/customer_support_comments/properties/#schema-created_until_timestamp) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/customer_support_comments/properties/#schema-name) |

## Next pages

- [comments](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/customer_support_comments/properties/comments/)
- [xcsh_customer_support_comments](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/customer_support_comments/)
