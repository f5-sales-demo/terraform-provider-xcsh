---
page_title: "xcsh_access_active_sessions reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_access_active_sessions reference."
---

# xcsh_access_active_sessions reference

<a id="canonical-ea3f8a31392a8fec51cb6b0a32c4892b70baf6ac0574b3877fc23d6b78865c9b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d34662e47616152eaa63ac6c036e6e7a240ccde09023111060b5563fc244f6ed"></a>

## Property reference — Property reference / dde9fb2df1e9 / 2

Breadcrumbs:

- [xcsh_access_active_sessions](../data-sources/access_active_sessions.md#canonical-d80936c849771550fc4fe0718c8a494dc68ce51087b0c6d03d5c8f5446ecad9f)
- Property reference

<a id="canonical-ce968c7a6400bb4eba59678f5d8116429c74da930e273e80950b442abbf33b4f"></a>

## Direct properties — Property reference / dde9fb2df1e9 / 3

<a id="canonical-12a557d6cccd57e975b24d7aea6cebdb42e3eeaf72c7a783fa21f4acc52c7422"></a>

<a id="canonical-50e5c4eb51f431f4b8a38ed97fb0a01823f28e1fe28524ab853fa60fcc7a9656"></a>

## client_ip property — Property reference / dde9fb2df1e9 / 4

Type: `"string"`. Optional.

Filter sessions by client IP address. Supports a comma-separated list to match any of the specified
values (e.g. '192.0.2.106,192.0.2.233').

<a id="canonical-8645bad55f12a163ee06d34f0c2974b2e7948352380d5b5b7b96b3509de7602f"></a>

<a id="canonical-c70ef20c514ef36ba616e568f2a2bab890e29fbe1adaf1458ef3899769417383"></a>

## cursor property — Property reference / dde9fb2df1e9 / 5

Type: `"string"`. Optional.

Cursor for pagination. Contains encoded sequence\_id for next/previous page.

<a id="canonical-2801d025b870457e8677f7fa316d2357fa8c539e9c697e8969237b379fec05c3"></a>

<a id="canonical-08f338017b25d33fcbdf9153e657255b266b2c0131dcf50bf8175b7c4397e627"></a>

## direction property — Property reference / dde9fb2df1e9 / 6

Type: `"string"`. Optional.

\[Enum: FORWARD|BACKWARD\] Direction for pagination (forward or backward) GET next page of results
GET previous page of results. Possible values are \`FORWARD\`, \`BACKWARD\`. Defaults to
\`FORWARD\`.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("FORWARD",
    "BACKWARD"),
}
```

<a id="canonical-f2b9d2f678abda8c2f108d3529f143c5522c22381102b0f38db697a04677c340"></a>

<a id="canonical-366e73a1c9784b8d51c3a4e3fd1517f7fe902a7b02dbe92514e73f0233f59041"></a>

## expiration_time_from property — Property reference / dde9fb2df1e9 / 7

Type: `"string"`. Optional.

Filter sessions that expire from this timestamp (inclusive) Format: RFC 3339 (e.g.,
2024-08-06T20:00:00Z)

<a id="canonical-5355f34643be21822550d5f14286b7aeb1b9390e3b292b7343f0375711d7b1ba"></a>

<a id="canonical-dcb5cbdc1ff25e16be692078e83b6c6a06f1e63395562880aa9a1a3cfb9b4439"></a>

## expiration_time_to property — Property reference / dde9fb2df1e9 / 8

Type: `"string"`. Optional.

Filter sessions that expire up to this timestamp (inclusive) Format: RFC 3339 (e.g.,
2024-08-07T20:00:00Z)

<a id="canonical-78ee7ae5d1c8e2d54d440e4ea99b54815d6e6b6d3a8803d07e320c820fe654ac"></a>

<a id="canonical-4646dd2ebed2eb5561092e40411485b82003d5b33d35fd0c27f87e745eae56e8"></a>

## item_count property — Property reference / dde9fb2df1e9 / 9

Type: `"number"`. Computed.

Total count of all active sessions (across all pages, after applying filters).

- [items](data-sources--access_active_sessions--reference--group-001.md#canonical-73b9f2a4d03e96ef2217fa7deb3ee0466754959949ee0db58d9fee1345b88d7a): complete subsection reference.

<a id="canonical-15dcfe31b11672ce9f3b351b03b3f939b166ac0360c0dd1b4e13417e0550fb1f"></a>

<a id="canonical-7c96a0f41c474ab1314fcf377d8833ecc942725b88b5f32ed5ebf79dd6cff843"></a>

## limit property — Property reference / dde9fb2df1e9 / 10

Type: `"number"`. Optional.

Limits the number of results to the specified number.

<a id="canonical-a7d5daf98b8846fa1b8d01ecf3422f1b588dbd1f8e2c0def4d2d0ad79e289b88"></a>

<a id="canonical-0190e375a311911fac25e842561891853687e840980462aa08b12bb68bb1c0c2"></a>

## namespace property — Property reference / dde9fb2df1e9 / 11

Type: `"string"`. Required.

Namespace Namespace of the App type for the current request.

<a id="canonical-29e6f5a897fd18a0bcfd2cbbd590e084f92366ea39727ef81febf9702d0924c3"></a>

<a id="canonical-80d4f3ea97f91ebc0afc4d0abfd3cb409e3bc9c75c405e5d2f1f4d6152679c0a"></a>

## next_cursor property — Property reference / dde9fb2df1e9 / 12

Type: `"string"`. Computed.

Next Cursor. Cursor for the next page of results.

<a id="canonical-e2a61b58402d1de72aef33ee3696a45dbb625755d5813b60dd37bec693d4f444"></a>

<a id="canonical-784dde6ecec063a53dd24cd78635087f4a84e5120d6a7b552ae76c4b04388107"></a>

## policy property — Property reference / dde9fb2df1e9 / 13

Type: `"string"`. Optional.

Filter sessions by policy. Supports a comma-separated list to match any of the specified values
(e.g. 'default-policy,strict-policy').

<a id="canonical-f08d76506cbb908ebfc686817919f1054a714f5b59af07b3d51349d41694c294"></a>

<a id="canonical-4037ee56b6045bef0231363a9aa071daba0f745207128f67b6a8a7ef89bc5b80"></a>

## previous_cursor property — Property reference / dde9fb2df1e9 / 14

Type: `"string"`. Computed.

Cursor for the previous page of results.

<a id="canonical-42e493439ba40d93a3105e20e39d5cd1fa75f976a27c6ceacfb32b750d84d17e"></a>

<a id="canonical-8bb0b2909a81c85eac5474ddce104ade6ac8c43130328ca4c78b26b523c9a779"></a>

## site property — Property reference / dde9fb2df1e9 / 15

Type: `"string"`. Optional.

Filter sessions by site. Supports a comma-separated list to match any of the specified values (e.g.
'us-west-2,eu-west-1').

<a id="canonical-7ed134536041f8f0da2a3edfe4f39d51cc2b37f8ed493a671af38a82f18bfdef"></a>

<a id="canonical-66ef60dd91951447c631a2ed2123f5cb11de4c81a6adc6cce85952f8f65f7864"></a>

## start_time_from property — Property reference / dde9fb2df1e9 / 16

Type: `"string"`. Optional.

Filter sessions that started from this timestamp (inclusive) Format: RFC 3339 (e.g.,
2024-08-06T20:00:00Z)

<a id="canonical-929e96f4c6a067bf4afef55900599b3907cfdf0ad0dd15ae65b82027d5d2ec7d"></a>

<a id="canonical-2e3d0064d967e10b81327ee726e6525690ceedd5847141bfa55eb87034738e92"></a>

## start_time_to property — Property reference / dde9fb2df1e9 / 17

Type: `"string"`. Optional.

Filter sessions that started up to this timestamp (inclusive) Format: RFC 3339 (e.g.,
2024-08-07T20:00:00Z)

<a id="canonical-82c471ba37bdb029b56c1998a68e093317e2c5521acf989794d3f2eb49668bfd"></a>

<a id="canonical-01a90e101d5a2f20641d91411ed16146fe45ab18cf1e3e9f335a80b6461770d3"></a>

## status property — Property reference / dde9fb2df1e9 / 18

Type: `"string"`. Optional.

\[Enum: ANY|PENDING\_ONLY|ESTABLISHED\_ONLY\] Filter sessions by status. If not specified, returns
sessions with all statuses. Return sessions with any status (default when no status filter is
specified) Filter for sessions with pending status Filter for sessions with established status.
Possible values are \`ANY\`, \`PENDING\_ONLY\`, \`ESTABLISHED\_ONLY\`. Defaults to \`ANY\`.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("ANY",
    "PENDING_ONLY",
    "ESTABLISHED_ONLY"),
}
```

<a id="canonical-3d410045e793e826567ded2ccca0d89a9f0c41835b4eb165566a62ad28c0753a"></a>

<a id="canonical-36a925040c83eb943b0543e6cc05c890c9c2dbccea8c8a6602833f55ae5bca07"></a>

## total_established property — Property reference / dde9fb2df1e9 / 19

Type: `"number"`. Computed.

Total count of sessions with ESTABLISHED status (after applying filters).

<a id="canonical-3d75e2a3284a40ef51481f705048cd4f8d963c1df472f5bf4e759d0ccc35fad3"></a>

<a id="canonical-14e26938b08876c2c9468daa5edef98e4de0376aba4f7c150b6c120d940567bc"></a>

## total_pending property — Property reference / dde9fb2df1e9 / 20

Type: `"number"`. Computed.

Total count of sessions with PENDING status (after applying filters).

<a id="canonical-31a515efc72620416de9102d8567b9bd4e7d646b386d898e2437e6c4d0b3dfbc"></a>

<a id="canonical-e3ea6749ad9243afa8e31f2b8aca8e8c0fcbc668a1d9f095ecadceda63dab7c1"></a>

## username property — Property reference / dde9fb2df1e9 / 21

Type: `"string"`. Optional.

Filter sessions by username. Supports a comma-separated list to match any of the specified values
(e.g. 'joe,bob').

<a id="canonical-4c436033e63710624fcd50536a0e89928031e490471d445f0d42ae00978d027e"></a>

<a id="canonical-77011a4ba42e85ad3c3adf032bf0d556bccaebe9d8b3d76597769e19269fcc28"></a>

## virtual_server property — Property reference / dde9fb2df1e9 / 22

Type: `"string"`. Optional.

Filter sessions by virtual server. Supports a comma-separated list to match any of the specified
values (e.g. 'web-app,API-server').

<a id="canonical-7c1157f5d542525fcfead51d202fb65897e7cae6175fbee5114f32e629bb7dfd"></a>

## All schema paths — Property reference / dde9fb2df1e9 / 23

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `client_ip` | [client_ip](data-sources--access_active_sessions--reference--group-001.md#canonical-12a557d6cccd57e975b24d7aea6cebdb42e3eeaf72c7a783fa21f4acc52c7422) |
| `cursor` | [cursor](data-sources--access_active_sessions--reference--group-001.md#canonical-8645bad55f12a163ee06d34f0c2974b2e7948352380d5b5b7b96b3509de7602f) |
| `direction` | [direction](data-sources--access_active_sessions--reference--group-001.md#canonical-2801d025b870457e8677f7fa316d2357fa8c539e9c697e8969237b379fec05c3) |
| `expiration_time_from` | [expiration_time_from](data-sources--access_active_sessions--reference--group-001.md#canonical-f2b9d2f678abda8c2f108d3529f143c5522c22381102b0f38db697a04677c340) |
| `expiration_time_to` | [expiration_time_to](data-sources--access_active_sessions--reference--group-001.md#canonical-5355f34643be21822550d5f14286b7aeb1b9390e3b292b7343f0375711d7b1ba) |
| `item_count` | [item_count](data-sources--access_active_sessions--reference--group-001.md#canonical-78ee7ae5d1c8e2d54d440e4ea99b54815d6e6b6d3a8803d07e320c820fe654ac) |
| `items` | [items](data-sources--access_active_sessions--reference--group-001.md#canonical-78c941ff3bb5daff5f56db8376aa8636b050ff28d1ba1cbb588200f11da6ef63) |
| `items.client_ip` | [items.client_ip](data-sources--access_active_sessions--reference--group-001.md#canonical-02b14d75b315ea9a0125ba3adb0fb43bcb0c64650bac4e149e1b61bc9e2209d1) |
| `items.expiration_time` | [items.expiration_time](data-sources--access_active_sessions--reference--group-001.md#canonical-d2ac313e977ad1eebb1b2bc5e9f29a1b8a0d64151f8ba29b8d25d4c94f404732) |
| `items.id` | [items.id](data-sources--access_active_sessions--reference--group-001.md#canonical-b0957d6873a57ab0533f0cb2e9a90699e582463cd4a726346f3cb2ca73a1ded5) |
| `items.last_activity_time` | [items.last_activity_time](data-sources--access_active_sessions--reference--group-001.md#canonical-dc2d013750e633c9151d0963d4f7db333f39a6b7cf33d421090588bd8a20c625) |
| `items.policy` | [items.policy](data-sources--access_active_sessions--reference--group-001.md#canonical-479c256150bcf99b169da33ac780fa40e2437eee936c31f76f99e8feffe10f9a) |
| `items.site` | [items.site](data-sources--access_active_sessions--reference--group-001.md#canonical-c88bc542695c439969bda9dfb23c9043096c1ea66eade9e01cb6b2ea24d874b9) |
| `items.start_time` | [items.start_time](data-sources--access_active_sessions--reference--group-001.md#canonical-618254c836a56b8becee97430fa18240e22f1a032afcc461f822fed45892a240) |
| `items.status` | [items.status](data-sources--access_active_sessions--reference--group-001.md#canonical-b4aa423a69c1372f0278d038a3c53d9f0f9d70ef973976885495b7bcfa23b486) |
| `items.username` | [items.username](data-sources--access_active_sessions--reference--group-001.md#canonical-8d48f4b9542fedbe1859560c8175b2e476bd7a0aa0f0bff283c7814674a897c6) |
| `items.virtual_server` | [items.virtual_server](data-sources--access_active_sessions--reference--group-001.md#canonical-69c9e5fce7b447fca48fc83df063d4aafcf6fc2064368dc5f84f36c0b237cfb4) |
| `limit` | [limit](data-sources--access_active_sessions--reference--group-001.md#canonical-15dcfe31b11672ce9f3b351b03b3f939b166ac0360c0dd1b4e13417e0550fb1f) |
| `namespace` | [namespace](data-sources--access_active_sessions--reference--group-001.md#canonical-a7d5daf98b8846fa1b8d01ecf3422f1b588dbd1f8e2c0def4d2d0ad79e289b88) |
| `next_cursor` | [next_cursor](data-sources--access_active_sessions--reference--group-001.md#canonical-29e6f5a897fd18a0bcfd2cbbd590e084f92366ea39727ef81febf9702d0924c3) |
| `policy` | [policy](data-sources--access_active_sessions--reference--group-001.md#canonical-e2a61b58402d1de72aef33ee3696a45dbb625755d5813b60dd37bec693d4f444) |
| `previous_cursor` | [previous_cursor](data-sources--access_active_sessions--reference--group-001.md#canonical-f08d76506cbb908ebfc686817919f1054a714f5b59af07b3d51349d41694c294) |
| `site` | [site](data-sources--access_active_sessions--reference--group-001.md#canonical-42e493439ba40d93a3105e20e39d5cd1fa75f976a27c6ceacfb32b750d84d17e) |
| `start_time_from` | [start_time_from](data-sources--access_active_sessions--reference--group-001.md#canonical-7ed134536041f8f0da2a3edfe4f39d51cc2b37f8ed493a671af38a82f18bfdef) |
| `start_time_to` | [start_time_to](data-sources--access_active_sessions--reference--group-001.md#canonical-929e96f4c6a067bf4afef55900599b3907cfdf0ad0dd15ae65b82027d5d2ec7d) |
| `status` | [status](data-sources--access_active_sessions--reference--group-001.md#canonical-82c471ba37bdb029b56c1998a68e093317e2c5521acf989794d3f2eb49668bfd) |
| `total_established` | [total_established](data-sources--access_active_sessions--reference--group-001.md#canonical-3d410045e793e826567ded2ccca0d89a9f0c41835b4eb165566a62ad28c0753a) |
| `total_pending` | [total_pending](data-sources--access_active_sessions--reference--group-001.md#canonical-3d75e2a3284a40ef51481f705048cd4f8d963c1df472f5bf4e759d0ccc35fad3) |
| `username` | [username](data-sources--access_active_sessions--reference--group-001.md#canonical-31a515efc72620416de9102d8567b9bd4e7d646b386d898e2437e6c4d0b3dfbc) |
| `virtual_server` | [virtual_server](data-sources--access_active_sessions--reference--group-001.md#canonical-4c436033e63710624fcd50536a0e89928031e490471d445f0d42ae00978d027e) |

<a id="canonical-0f59b99265a1d8966487136277e2db2dd86ddfc27fab5db5f7e44f94aab83460"></a>

## Next pages — Property reference / dde9fb2df1e9 / 24

- [items](data-sources--access_active_sessions--reference--group-001.md#canonical-73b9f2a4d03e96ef2217fa7deb3ee0466754959949ee0db58d9fee1345b88d7a)
- [xcsh_access_active_sessions](../data-sources/access_active_sessions.md#canonical-d80936c849771550fc4fe0718c8a494dc68ce51087b0c6d03d5c8f5446ecad9f)

<a id="canonical-73b9f2a4d03e96ef2217fa7deb3ee0466754959949ee0db58d9fee1345b88d7a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-36cd7b733538383531e47b6e5d9a3f5f6b599a695d9013c22556409592df890d"></a>

## items — items / 46eabbb0004a / 2

Breadcrumbs:

- [xcsh_access_active_sessions](../data-sources/access_active_sessions.md#canonical-d80936c849771550fc4fe0718c8a494dc68ce51087b0c6d03d5c8f5446ecad9f)
- [Property reference](data-sources--access_active_sessions--reference--group-001.md#canonical-ea3f8a31392a8fec51cb6b0a32c4892b70baf6ac0574b3877fc23d6b78865c9b)
- items

<a id="canonical-78c941ff3bb5daff5f56db8376aa8636b050ff28d1ba1cbb588200f11da6ef63"></a>

Type: `"list"`. Computed.

Sessions. List of active sessions.

<a id="canonical-6d2b305c108bf2668652f074d698ed490632d31faee6c7781d454cdd815fc477"></a>

## Direct properties — items / 46eabbb0004a / 3

<a id="canonical-02b14d75b315ea9a0125ba3adb0fb43bcb0c64650bac4e149e1b61bc9e2209d1"></a>

<a id="canonical-11e88fbfab5cc4e3c983a29ca97346f3ab57ec1cac38893f3dd0fdae89fd9361"></a>

## client_ip property — items / 46eabbb0004a / 4

Type: `"string"`. Computed.

Client IP of the user that connected via the session.

<a id="canonical-d2ac313e977ad1eebb1b2bc5e9f29a1b8a0d64151f8ba29b8d25d4c94f404732"></a>

<a id="canonical-9ee66a0461c02d0579b8fe789020ebb96acea27d0881b88e0e77ced9c105d10d"></a>

## expiration_time property — items / 46eabbb0004a / 5

Type: `"string"`. Computed.

Expiration time of the session in RFC 3339 format.

<a id="canonical-b0957d6873a57ab0533f0cb2e9a90699e582463cd4a726346f3cb2ca73a1ded5"></a>

<a id="canonical-c1369a06d029568f2f9171858615eabdac95967079d1132b5289f35058235e41"></a>

## id property — items / 46eabbb0004a / 6

Type: `"string"`. Computed.

Session ID. ID of the session.

<a id="canonical-dc2d013750e633c9151d0963d4f7db333f39a6b7cf33d421090588bd8a20c625"></a>

<a id="canonical-635cc938f554a46275d727adf183a0be859376d8111dd01a13152f16ea243e94"></a>

## last_activity_time property — items / 46eabbb0004a / 7

Type: `"string"`. Computed.

Last activity time of the session in RFC 3339 format.

<a id="canonical-479c256150bcf99b169da33ac780fa40e2437eee936c31f76f99e8feffe10f9a"></a>

<a id="canonical-ad35d78741326b8ecabae1f46b171c67f3fd8cfb0f08e95d0cb13f8ccc5926d2"></a>

## policy property — items / 46eabbb0004a / 8

Type: `"string"`. Computed.

Policy. Policy associated with the session.

<a id="canonical-c88bc542695c439969bda9dfb23c9043096c1ea66eade9e01cb6b2ea24d874b9"></a>

<a id="canonical-ed671951e0ce7c1d14c63d013ebef7738a0b122a6bd7a3054afc6535b87f0387"></a>

## site property — items / 46eabbb0004a / 9

Type: `"string"`. Computed.

Site. Site where the session is created.

<a id="canonical-618254c836a56b8becee97430fa18240e22f1a032afcc461f822fed45892a240"></a>

<a id="canonical-238721038686437942ee8b162fef080ca5be885fb49cde9a6ddbf4e727daab79"></a>

## start_time property — items / 46eabbb0004a / 10

Type: `"string"`. Computed.

Start time of the session in RFC 3339 format.

<a id="canonical-b4aa423a69c1372f0278d038a3c53d9f0f9d70ef973976885495b7bcfa23b486"></a>

<a id="canonical-b19d823348c8eeb1a94021c2a25e2b4c85287d0d73fe4d438adff598addf1d5b"></a>

## status property — items / 46eabbb0004a / 11

Type: `"string"`. Computed.

\[Enum: PENDING|ESTABLISHED\] The actual status of an active session Session status is pending
Session status is established. Possible values are \`PENDING\`, \`ESTABLISHED\`. Defaults to
\`PENDING\`.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("PENDING",
    "ESTABLISHED"),
}
```

<a id="canonical-8d48f4b9542fedbe1859560c8175b2e476bd7a0aa0f0bff283c7814674a897c6"></a>

<a id="canonical-d7b2715f2b16544b7af5fdd7c8a451ccf594da854265a467df81ec063597475a"></a>

## username property — items / 46eabbb0004a / 12

Type: `"string"`. Computed.

Username used in authentication of this session.

<a id="canonical-69c9e5fce7b447fca48fc83df063d4aafcf6fc2064368dc5f84f36c0b237cfb4"></a>

<a id="canonical-7212d617eca294c70558e1a5242335f0bafd19bda90f833d0d58669046296ea5"></a>

## virtual_server property — items / 46eabbb0004a / 13

Type: `"string"`. Computed.

Virtual server associated with the session.

<a id="canonical-ef7148966795a39d428df81a00802c9fa30ff1ff337a3d19ccc7af51a9144a29"></a>

## Next pages — items / 46eabbb0004a / 14

- [Property reference](data-sources--access_active_sessions--reference--group-001.md#canonical-ea3f8a31392a8fec51cb6b0a32c4892b70baf6ac0574b3877fc23d6b78865c9b)
- [xcsh_access_active_sessions](../data-sources/access_active_sessions.md#canonical-d80936c849771550fc4fe0718c8a494dc68ce51087b0c6d03d5c8f5446ecad9f)
