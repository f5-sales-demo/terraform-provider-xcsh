---
page_title: "xcsh_smsv2_aws_runtime reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_smsv2_aws_runtime reference."
---

# xcsh_smsv2_aws_runtime reference

<a id="canonical-5d8cf40e83f05228ca10a7a18a1f29237fffea18c626c7b13a3340addaede44e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9415693931adba64039d0a41f01062e5c36a8f2897e6b289050be84709bcd385"></a>

## Property reference — Property reference / 7ca3175b3718 / 2

Breadcrumbs:

- [xcsh_smsv2_aws_runtime](../data-sources/smsv2_aws_runtime.md#canonical-585eec57272cb79858684d22159d08a11b9cc9e8b386856fa87be603c746a674)
- Property reference

<a id="canonical-123612d81c981af78b07ff214da432b56a5d2d400eaf6312536bd0af9700b977"></a>

## Direct properties — Property reference / 7ca3175b3718 / 3

<a id="canonical-e822d4ae2ee4d744cdeef08eda2bf902413266b6ca63d08b8862dfdc07118b47"></a>

<a id="canonical-5bcd3712c4ecb6ee4b5896bfc6df4e6aabcf3372cf657ad584875beab84adfb4"></a>

## healthy property — Property reference / 7ca3175b3718 / 4

Type: `"bool"`. Computed.

<a id="canonical-92730d95aa07f41d8e81b96ea4d5d6115932910fc884334749598223018a9715"></a>

<a id="canonical-9aa9872379a44ce67116df5363be7958fd782486f892364b4b624f415510857b"></a>

## id property — Property reference / 7ca3175b3718 / 5

Type: `"string"`. Computed.

- [interfaces](data-sources--smsv2_aws_runtime--reference--group-001.md#canonical-245af834de489b8dd6ffedab74132879ab2adcd2f524d627c02e39dbe62d28a8): complete subsection reference.

<a id="canonical-f91d16ae41836fdea9da48f565bc7a9881701ac7609c52e14b09794b38994e61"></a>

<a id="canonical-de596107fae250b8b3a9abdbe1fdff8b902de3c695a968a317c8c12df68c9639"></a>

## namespace property — Property reference / 7ca3175b3718 / 6

Type: `"string"`. Required.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{stringvalidator.OneOf("system")}
```

- [nodes](data-sources--smsv2_aws_runtime--reference--group-001.md#canonical-d0339c0f31fa8eeac57f1a279e07954b5d9299b8c504997e3abcc668418dda1e): complete subsection reference.

<a id="canonical-601a5a47a854c121b3b278649efd105d262617569950551a9c706a32cf80a667"></a>

<a id="canonical-d47882fb26facd37d7380c0b20225a45363806633c8d7bda710069aa3814449a"></a>

## poll_interval_seconds property — Property reference / 7ca3175b3718 / 7

Type: `"number"`. Optional, Computed.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{int64validator.Between(1, 60)}
```

<a id="canonical-2edea3b355b8b88a86c31e0169f402b855250a6dd7a2e16055182e78e1dc912b"></a>

<a id="canonical-ad240797612b4750c49eab80eba4daa84cba2e3c243162eceaaea0ecb42123c9"></a>

## site property — Property reference / 7ca3175b3718 / 8

Type: `"string"`. Required.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{stringvalidator.LengthAtLeast(1)}
```

<a id="canonical-0ab0a54aa28442d534abba3aa2b9d423026954ec6cc9a9b34af2cce21554b3b5"></a>

<a id="canonical-3f920341480e4743b50440b9f866e6accd0577546cb285e680cada98f3d09043"></a>

## timeout_seconds property — Property reference / 7ca3175b3718 / 9

Type: `"number"`. Optional, Computed.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{int64validator.Between(1, 7200)}
```

<a id="canonical-b1a3cb5b0277c9ea2de1ce91ece7842937c514b462a51ecf7950088b0f0badaf"></a>

## All schema paths — Property reference / 7ca3175b3718 / 10

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `healthy` | [healthy](data-sources--smsv2_aws_runtime--reference--group-001.md#canonical-e822d4ae2ee4d744cdeef08eda2bf902413266b6ca63d08b8862dfdc07118b47) |
| `id` | [id](data-sources--smsv2_aws_runtime--reference--group-001.md#canonical-92730d95aa07f41d8e81b96ea4d5d6115932910fc884334749598223018a9715) |
| `interfaces` | [interfaces](data-sources--smsv2_aws_runtime--reference--group-001.md#canonical-621faf2e063d86aed0e9376d1641d6ff665a944f86d9a5da9e40c592724237c1) |
| `interfaces.healthy` | [interfaces.healthy](data-sources--smsv2_aws_runtime--reference--group-001.md#canonical-d865cb06e6fa786b28582d9f7772b2b8f8e7f8f36c87047414d335f579c4b95b) |
| `interfaces.interface_name` | [interfaces.interface_name](data-sources--smsv2_aws_runtime--reference--group-001.md#canonical-5c86688419140ef7fb6d3467d05801d93709efd186f3b4ee4bcedfac4006cc1d) |
| `interfaces.mac` | [interfaces.mac](data-sources--smsv2_aws_runtime--reference--group-001.md#canonical-bee19bf98563e45879bee82d2341a17a97374954db2dc5ffc7909def76dba7cc) |
| `interfaces.mtu` | [interfaces.mtu](data-sources--smsv2_aws_runtime--reference--group-001.md#canonical-4948245275ba7f0aef9ca1850eb6a5ff8591fa6cd3a7045b035827159145e8f4) |
| `interfaces.node` | [interfaces.node](data-sources--smsv2_aws_runtime--reference--group-001.md#canonical-ffaa8e439953ef0b32f2d3b18ecd64bc1e830a5d8bbaca3de888948e81a1d3cb) |
| `interfaces.role` | [interfaces.role](data-sources--smsv2_aws_runtime--reference--group-001.md#canonical-917a073b38b068e2ada5e34232bfa743c5cbd9c4139fd0080933f8e11a8d5647) |
| `namespace` | [namespace](data-sources--smsv2_aws_runtime--reference--group-001.md#canonical-f91d16ae41836fdea9da48f565bc7a9881701ac7609c52e14b09794b38994e61) |
| `nodes` | [nodes](data-sources--smsv2_aws_runtime--reference--group-001.md#canonical-9c115e3a27b0920d2ec71eccc0494280ffd91d86e51c996cc0ea807186bc28fb) |
| `nodes.mac` | [nodes.mac](data-sources--smsv2_aws_runtime--reference--group-001.md#canonical-d19aa10a8b4a1fd908f92ee73009ac7de66ebeec987b015f599604a446ea2a79) |
| `nodes.node` | [nodes.node](data-sources--smsv2_aws_runtime--reference--group-001.md#canonical-97ddc04226c9c675be7527ffc1282960f552f0a913277c62b7ab055eb6e6301a) |
| `nodes.role` | [nodes.role](data-sources--smsv2_aws_runtime--reference--group-001.md#canonical-79559f790e38ac6463423cb25fd7aaab204b17e61ad8eb272df50ceff3638bff) |
| `poll_interval_seconds` | [poll_interval_seconds](data-sources--smsv2_aws_runtime--reference--group-001.md#canonical-601a5a47a854c121b3b278649efd105d262617569950551a9c706a32cf80a667) |
| `site` | [site](data-sources--smsv2_aws_runtime--reference--group-001.md#canonical-2edea3b355b8b88a86c31e0169f402b855250a6dd7a2e16055182e78e1dc912b) |
| `timeout_seconds` | [timeout_seconds](data-sources--smsv2_aws_runtime--reference--group-001.md#canonical-0ab0a54aa28442d534abba3aa2b9d423026954ec6cc9a9b34af2cce21554b3b5) |

<a id="canonical-add968706c81cc99403b2470a9a690fcc77dc2ceccee18c98027764b083c8400"></a>

## Next pages — Property reference / 7ca3175b3718 / 11

- [interfaces](data-sources--smsv2_aws_runtime--reference--group-001.md#canonical-245af834de489b8dd6ffedab74132879ab2adcd2f524d627c02e39dbe62d28a8)
- [nodes](data-sources--smsv2_aws_runtime--reference--group-001.md#canonical-d0339c0f31fa8eeac57f1a279e07954b5d9299b8c504997e3abcc668418dda1e)
- [xcsh_smsv2_aws_runtime](../data-sources/smsv2_aws_runtime.md#canonical-585eec57272cb79858684d22159d08a11b9cc9e8b386856fa87be603c746a674)

<a id="canonical-245af834de489b8dd6ffedab74132879ab2adcd2f524d627c02e39dbe62d28a8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5265e9ac980bf994db054170a5b20e8dc0d85a52a884d82674d11bad607dc71f"></a>

## interfaces — interfaces / e2cb15918d73 / 2

Breadcrumbs:

- [xcsh_smsv2_aws_runtime](../data-sources/smsv2_aws_runtime.md#canonical-585eec57272cb79858684d22159d08a11b9cc9e8b386856fa87be603c746a674)
- [Property reference](data-sources--smsv2_aws_runtime--reference--group-001.md#canonical-5d8cf40e83f05228ca10a7a18a1f29237fffea18c626c7b13a3340addaede44e)
- interfaces

<a id="canonical-621faf2e063d86aed0e9376d1641d6ff665a944f86d9a5da9e40c592724237c1"></a>

Type: `"map"`. Computed.

<a id="canonical-5a3ce20aedb8c2ad6c49127f0831d627fc1260ba82085e5688d4e2d8b833345c"></a>

## Direct properties — interfaces / e2cb15918d73 / 3

<a id="canonical-d865cb06e6fa786b28582d9f7772b2b8f8e7f8f36c87047414d335f579c4b95b"></a>

<a id="canonical-d45226044ecb21828c02bde97f76d04aed5ad777fe45563ba1f02b06d94bc092"></a>

## healthy property — interfaces / e2cb15918d73 / 4

Type: `"bool"`. Computed.

<a id="canonical-5c86688419140ef7fb6d3467d05801d93709efd186f3b4ee4bcedfac4006cc1d"></a>

<a id="canonical-17913db7d930869377561d1a3c49bc69504843112ecd3d3554c3a8dfd2c97d40"></a>

## interface_name property — interfaces / e2cb15918d73 / 5

Type: `"string"`. Computed.

<a id="canonical-bee19bf98563e45879bee82d2341a17a97374954db2dc5ffc7909def76dba7cc"></a>

<a id="canonical-538869dd87ea86ba4bcf824c3c801e6b2c050d014b95d7a4e8972fc6b938ee6b"></a>

## mac property — interfaces / e2cb15918d73 / 6

Type: `"string"`. Computed.

<a id="canonical-4948245275ba7f0aef9ca1850eb6a5ff8591fa6cd3a7045b035827159145e8f4"></a>

<a id="canonical-dc011c77ef69ec4d77381c2db4a2c8489c4454173d70a1c10436216bba4b7881"></a>

## mtu property — interfaces / e2cb15918d73 / 7

Type: `"number"`. Computed.

<a id="canonical-ffaa8e439953ef0b32f2d3b18ecd64bc1e830a5d8bbaca3de888948e81a1d3cb"></a>

<a id="canonical-1e1e49e3cb6daa0fe7e910dc982acf220c12d0c4590b44703dbe6f57ee2b0c74"></a>

## node property — interfaces / e2cb15918d73 / 8

Type: `"string"`. Computed.

<a id="canonical-917a073b38b068e2ada5e34232bfa743c5cbd9c4139fd0080933f8e11a8d5647"></a>

<a id="canonical-ef87740b0f2027905680d3af12ba49601b70942b4137d9a980bfed2926c2a692"></a>

## role property — interfaces / e2cb15918d73 / 9

Type: `"string"`. Computed.

<a id="canonical-67ae618751b4bea1edc8ac281e751a87af585933dc4a4cf82024af49adbf9a4e"></a>

## Next pages — interfaces / e2cb15918d73 / 10

- [Property reference](data-sources--smsv2_aws_runtime--reference--group-001.md#canonical-5d8cf40e83f05228ca10a7a18a1f29237fffea18c626c7b13a3340addaede44e)
- [xcsh_smsv2_aws_runtime](../data-sources/smsv2_aws_runtime.md#canonical-585eec57272cb79858684d22159d08a11b9cc9e8b386856fa87be603c746a674)

<a id="canonical-d0339c0f31fa8eeac57f1a279e07954b5d9299b8c504997e3abcc668418dda1e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4d002635b128ebb999eee996c0ee24b3ac8798d7c63256fad4dc164b904cfc04"></a>

## nodes — nodes / dc1e4c9ec638 / 2

Breadcrumbs:

- [xcsh_smsv2_aws_runtime](../data-sources/smsv2_aws_runtime.md#canonical-585eec57272cb79858684d22159d08a11b9cc9e8b386856fa87be603c746a674)
- [Property reference](data-sources--smsv2_aws_runtime--reference--group-001.md#canonical-5d8cf40e83f05228ca10a7a18a1f29237fffea18c626c7b13a3340addaede44e)
- nodes

<a id="canonical-9c115e3a27b0920d2ec71eccc0494280ffd91d86e51c996cc0ea807186bc28fb"></a>

Type: `"map"`. Required.

<a id="canonical-7d6e93a9849b1d378c3649d909456829cb6424ad13a95a805ecf095888b0ce7a"></a>

## Direct properties — nodes / dc1e4c9ec638 / 3

<a id="canonical-d19aa10a8b4a1fd908f92ee73009ac7de66ebeec987b015f599604a446ea2a79"></a>

<a id="canonical-340a13f7db7a778e61889ac2643e90a92f2c8247c05d2e467abe5d0eb4b9906a"></a>

## mac property — nodes / dc1e4c9ec638 / 4

Type: `"string"`. Required.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{stringvalidator.LengthAtLeast(1)}
```

<a id="canonical-97ddc04226c9c675be7527ffc1282960f552f0a913277c62b7ab055eb6e6301a"></a>

<a id="canonical-541f5bf6316e27047e450c99e6bffb1a84e0622688e6a40b47233ffd9cb0d9af"></a>

## node property — nodes / dc1e4c9ec638 / 5

Type: `"string"`. Required.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{stringvalidator.LengthAtLeast(1)}
```

<a id="canonical-79559f790e38ac6463423cb25fd7aaab204b17e61ad8eb272df50ceff3638bff"></a>

<a id="canonical-c79427d8812767e874c5aff140ac3e88b1a50532bac733086b6ecb5fbbae68b6"></a>

## role property — nodes / dc1e4c9ec638 / 6

Type: `"string"`. Required.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{stringvalidator.OneOf("slo",
    "sli")}
```

<a id="canonical-86608c80fbf0a3b1ec33ce27f111b41eb423e9ae3958813f75cc081992ecc2fd"></a>

## Next pages — nodes / dc1e4c9ec638 / 7

- [Property reference](data-sources--smsv2_aws_runtime--reference--group-001.md#canonical-5d8cf40e83f05228ca10a7a18a1f29237fffea18c626c7b13a3340addaede44e)
- [xcsh_smsv2_aws_runtime](../data-sources/smsv2_aws_runtime.md#canonical-585eec57272cb79858684d22159d08a11b9cc9e8b386856fa87be603c746a674)
