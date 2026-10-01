---
page_title: "xcsh_smsv2_kvm_runtime reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_smsv2_kvm_runtime reference."
---

# xcsh_smsv2_kvm_runtime reference

<a id="canonical-c1611532e0fbc38a6575ad57eb82084c21632bc4f7f049080aa3ea62c86c7928"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b3d613fd5b0fa36f13354a8909cc20cbc7f65adddbb26ec29d8fec15c9273032"></a>

## Property reference — Property reference / 52d852a12224 / 2

Breadcrumbs:

- [xcsh_smsv2_kvm_runtime](../data-sources/smsv2_kvm_runtime.md#canonical-d6e6eff6744bfd966ef04ac1a8c4c3c58780b35501827bb7c51f1c87804071a2)
- Property reference

<a id="canonical-24fcdbec91a908924d8c040929a8e6607f16945fd14d50fd4770c953867ca0bb"></a>

## Direct properties — Property reference / 52d852a12224 / 3

<a id="canonical-24aefe4a8a8cd52cc03e21b12dd331ed49d32a8d51e5b6f78d7cdcf4556592b1"></a>

<a id="canonical-31d1cf65430adbf928ed9c11f52b8ccc6031ce273dc826cd2e2dd3285f7579d5"></a>

## device property — Property reference / 52d852a12224 / 4

Type: `"string"`. Computed.

Guest device reported by the matching live KVM registration.

<a id="canonical-930d770d8cee71dc15bcaa0e8a92ffd8eebe35ef94e01fe54fffdb70a3d86ae8"></a>

<a id="canonical-cd0ef79c75dc0932e6a8f12af7fb2ebc4f4c2a37750c5df104060cf1b548e931"></a>

## expected_mac property — Property reference / 52d852a12224 / 5

Type: `"string"`. Required.

Terraform-owned KVM CE MAC used to select one registration device.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{stringvalidator.LengthAtLeast(1)}
```

<a id="canonical-35f31b6bd306f57a37ca260371224168910f14f45fdbdd18f1af8dbeaae2844c"></a>

<a id="canonical-0ba581cef84be0f10df817515410eb02affd6262c1e2927e6540f775062b5474"></a>

## hostname property — Property reference / 52d852a12224 / 6

Type: `"string"`. Computed.

Hostname reported by the matching live KVM registration.

<a id="canonical-7e9212a21bcd2e1e2c68bdf34dc36854526895570ffbaa231dad255212dc6af5"></a>

<a id="canonical-a2536f981035937806478e6d3deb292a5fb5cebb2b2ae40920443069802835aa"></a>

## id property — Property reference / 52d852a12224 / 7

Type: `"string"`. Computed.

Stable lookup identity composed from the site and normalized expected MAC.

<a id="canonical-1893e08d3e8f5f9c97ebf3b44c204b22f28f9b1a75374379bc14be5668b63ffb"></a>

<a id="canonical-63a912a2f9db18cc460b21fcaadf4794a2911aa457b10cb54c852a034922a3cd"></a>

## interface_name property — Property reference / 52d852a12224 / 8

Type: `"string"`. Computed.

Exact owned XC network\_interface object name.

<a id="canonical-132e955d33a499c07907fede043e9ede577554839fe642b7dc832818585fb5fd"></a>

<a id="canonical-08ddf24a0b8e15952ff6e1bf1de400f7b41501e24f9c317ece36fde70feb11e6"></a>

## mac property — Property reference / 52d852a12224 / 9

Type: `"string"`. Computed.

Normalized six-octet MAC address.

<a id="canonical-cbc6c7e871655c52b1d734c15490da9266d8f1aac6bdc48b5b97018076badea5"></a>

<a id="canonical-981f3896ea6daa61cf760cf2bdadfa7594dc4ad74230e0a18bf374632c823183"></a>

## namespace property — Property reference / 52d852a12224 / 10

Type: `"string"`. Optional, Computed.

Namespace containing the site and realized interface. Defaults to \`system\`.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{stringvalidator.OneOf("system")}
```

<a id="canonical-e04823881fd152100aaceea5d667e28e07bafdc01909f9247ac1e89eceb9b28e"></a>

<a id="canonical-181b1eb3aca85a522ff9e5653a231723ef877677b083845e562b9e6cc4ea5c79"></a>

## online property — Property reference / 52d852a12224 / 11

Type: `"bool"`. Computed.

Whether the matching registration currently reports ONLINE.

<a id="canonical-82613fa425615aee7df243d71f515ca5079d035d6bf7c375587094a30f813a4d"></a>

<a id="canonical-00c53aefefa05eb38424b8cde197fa5e686905ec06b2700eedb78a0ef5b8de34"></a>

## poll_interval_seconds property — Property reference / 52d852a12224 / 12

Type: `"number"`. Optional, Computed.

Polling interval. Defaults to 10 seconds.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{int64validator.Between(1, 60)}
```

<a id="canonical-249672ad92f9791c87a2d8825ee5fe259b240db52dee5464e61369562dc99c55"></a>

<a id="canonical-0ecf68082ece1fd13a0ab03297851dd995f40b239dc0dfdb02a5fd06a5db14d8"></a>

## registration_state property — Property reference / 52d852a12224 / 13

Type: `"string"`. Computed.

Current state of the matching registration.

<a id="canonical-a784066618b4682ce942b0cae0db031a381bc0e484fb854644196b50225e1a5e"></a>

<a id="canonical-9fc03ce878f27dc4995ce966729b147e411bf90486d984f22e504567e6974af1"></a>

## site property — Property reference / 52d852a12224 / 14

Type: `"string"`. Required.

Secure Mesh Site v2 configuration name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{stringvalidator.LengthAtLeast(1)}
```

<a id="canonical-ac624ce702d82add6dfe378dc77ef386c232b2e3485f7612d7d6b0b3d0af2d0a"></a>

<a id="canonical-695038ab63ad5266032ca74e00cfc1364ed418125bfee1e93a5d40d86be2b0d0"></a>

## timeout_seconds property — Property reference / 52d852a12224 / 15

Type: `"number"`. Optional, Computed.

Bounded runtime discovery timeout. Defaults to 7200 seconds.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{int64validator.Between(1, 7200)}
```

<a id="canonical-8f44cb523c6197941017cfefb065956e4a02c2e3bfc3e8f72b971434a3641b22"></a>

## All schema paths — Property reference / 52d852a12224 / 16

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `device` | [device](data-sources--smsv2_kvm_runtime--reference--group-001.md#canonical-24aefe4a8a8cd52cc03e21b12dd331ed49d32a8d51e5b6f78d7cdcf4556592b1) |
| `expected_mac` | [expected_mac](data-sources--smsv2_kvm_runtime--reference--group-001.md#canonical-930d770d8cee71dc15bcaa0e8a92ffd8eebe35ef94e01fe54fffdb70a3d86ae8) |
| `hostname` | [hostname](data-sources--smsv2_kvm_runtime--reference--group-001.md#canonical-35f31b6bd306f57a37ca260371224168910f14f45fdbdd18f1af8dbeaae2844c) |
| `id` | [id](data-sources--smsv2_kvm_runtime--reference--group-001.md#canonical-7e9212a21bcd2e1e2c68bdf34dc36854526895570ffbaa231dad255212dc6af5) |
| `interface_name` | [interface_name](data-sources--smsv2_kvm_runtime--reference--group-001.md#canonical-1893e08d3e8f5f9c97ebf3b44c204b22f28f9b1a75374379bc14be5668b63ffb) |
| `mac` | [mac](data-sources--smsv2_kvm_runtime--reference--group-001.md#canonical-132e955d33a499c07907fede043e9ede577554839fe642b7dc832818585fb5fd) |
| `namespace` | [namespace](data-sources--smsv2_kvm_runtime--reference--group-001.md#canonical-cbc6c7e871655c52b1d734c15490da9266d8f1aac6bdc48b5b97018076badea5) |
| `online` | [online](data-sources--smsv2_kvm_runtime--reference--group-001.md#canonical-e04823881fd152100aaceea5d667e28e07bafdc01909f9247ac1e89eceb9b28e) |
| `poll_interval_seconds` | [poll_interval_seconds](data-sources--smsv2_kvm_runtime--reference--group-001.md#canonical-82613fa425615aee7df243d71f515ca5079d035d6bf7c375587094a30f813a4d) |
| `registration_state` | [registration_state](data-sources--smsv2_kvm_runtime--reference--group-001.md#canonical-249672ad92f9791c87a2d8825ee5fe259b240db52dee5464e61369562dc99c55) |
| `site` | [site](data-sources--smsv2_kvm_runtime--reference--group-001.md#canonical-a784066618b4682ce942b0cae0db031a381bc0e484fb854644196b50225e1a5e) |
| `timeout_seconds` | [timeout_seconds](data-sources--smsv2_kvm_runtime--reference--group-001.md#canonical-ac624ce702d82add6dfe378dc77ef386c232b2e3485f7612d7d6b0b3d0af2d0a) |

<a id="canonical-b86cb18afd305156e19d55f5c8001768899346a0343345c846d45aacb2fe4bd8"></a>

## Next pages — Property reference / 52d852a12224 / 17

- [xcsh_smsv2_kvm_runtime](../data-sources/smsv2_kvm_runtime.md#canonical-d6e6eff6744bfd966ef04ac1a8c4c3c58780b35501827bb7c51f1c87804071a2)
