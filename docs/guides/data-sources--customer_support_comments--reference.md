---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_customer_support_comments."
xcsh_docs: {"aliases": [], "body_bytes": 3729, "body_sha256": "sha256:3dccec49fa93b7e0c7c36dddc228f0fee81f5ceda871d6c2b8135a9109674236", "canonical_id": "xcsh-docs:data-sources:customer_support_comments:reference", "child_ids": ["xcsh-docs:data-sources:customer_support_comments:properties:comments"], "collection_id": "xcsh-docs:data-sources:customer_support_comments:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:customer_support_comments:reference", "parent_id": "xcsh-docs:data-sources:customer_support_comments:fundamentals", "path": "docs/guides/data-sources--customer_support_comments--reference.md", "provider_name": "customer_support_comments", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/customer_support_comments/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_customer_support_comments.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_customer_support_comments](../data-sources/customer_support_comments.md)
- Property reference

## Direct properties

<a id="schema-all_comments_returned"></a>

### all_comments_returned property

Type: `"bool"`. Computed.

Indicates if all comments for the issue have been returned.

- [comments](data-sources--customer_support_comments--properties--comments.md): complete subsection reference.

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
| `all_comments_returned` | [all_comments_returned](data-sources--customer_support_comments--reference.md#schema-all_comments_returned) |
| `comments` | [comments](data-sources--customer_support_comments--properties--comments.md#section) |
| `comments.attachment_ids` | [comments.attachment_ids](data-sources--customer_support_comments--properties--comments.md#schema-comments--attachment_ids) |
| `comments.attachments_info` | [comments.attachments_info](data-sources--customer_support_comments--properties--comments--attachments_info.md#section) |
| `comments.attachments_info.attachment` | [comments.attachments_info.attachment](data-sources--customer_support_comments--properties--comments--attachments_info.md#schema-comments--attachments_info--attachment) |
| `comments.attachments_info.content_type` | [comments.attachments_info.content_type](data-sources--customer_support_comments--properties--comments--attachments_info.md#schema-comments--attachments_info--content_type) |
| `comments.attachments_info.filename` | [comments.attachments_info.filename](data-sources--customer_support_comments--properties--comments--attachments_info.md#schema-comments--attachments_info--filename) |
| `comments.attachments_info.tp_id` | [comments.attachments_info.tp_id](data-sources--customer_support_comments--properties--comments--attachments_info.md#schema-comments--attachments_info--tp_id) |
| `comments.author_email` | [comments.author_email](data-sources--customer_support_comments--properties--comments.md#schema-comments--author_email) |
| `comments.author_name` | [comments.author_name](data-sources--customer_support_comments--properties--comments.md#schema-comments--author_name) |
| `comments.comment_id` | [comments.comment_id](data-sources--customer_support_comments--properties--comments.md#schema-comments--comment_id) |
| `comments.created_at` | [comments.created_at](data-sources--customer_support_comments--properties--comments.md#schema-comments--created_at) |
| `comments.html` | [comments.html](data-sources--customer_support_comments--properties--comments.md#schema-comments--html) |
| `comments.plain_text` | [comments.plain_text](data-sources--customer_support_comments--properties--comments.md#schema-comments--plain_text) |
| `created_until_timestamp` | [created_until_timestamp](data-sources--customer_support_comments--reference.md#schema-created_until_timestamp) |
| `name` | [name](data-sources--customer_support_comments--reference.md#schema-name) |

## Next pages

- [comments](data-sources--customer_support_comments--properties--comments.md)
- [xcsh_customer_support_comments](../data-sources/customer_support_comments.md)
