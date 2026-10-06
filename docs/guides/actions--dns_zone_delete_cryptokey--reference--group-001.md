---
page_title: "xcsh_dns_zone_delete_cryptokey reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_dns_zone_delete_cryptokey reference."
---

# xcsh_dns_zone_delete_cryptokey reference

<a id="canonical-1103032100322021-2001303320000110-2331221323002003-2030210003223133-1201102310333011-1330111231102102-1301022103211223-2321303012132222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_dns_zone_delete_cryptokey](../actions/dns_zone_delete_cryptokey.md#canonical-3100330101232133-3311113231220332-0022020131232233-2130012232123321-2121220121030020-0323223001111032-0003233110133020-2000133100010203)
- Property reference

<a id="canonical-0330302122223130-1302113112221320-1321203213223121-2100003311030233-0312020311221033-3032221133033021-2323032110031002-3132323131313013"></a>

### Direct properties for `xcsh_dns_zone_delete_cryptokey`

<a id="canonical-3301211003310303-0111331331031121-0220001010303130-2011230333323112-0231303310301032-2322203123103010-2232201101333130-1232310022130312"></a>

#### `key_id` property

Type: `"number"`. Optional.

Key ID. Unique identifier for this resource

<a id="canonical-3223321031030132-0133021021112212-0112002012020331-3222120003100311-2012020020022101-1023120300332102-2222303003121130-1213333031320003"></a>

<a id="canonical-0022121303002311-1220331131223101-1001110102212102-3101211201133220-1202330223330331-0113031310110111-3220221013012321-3221122322211011"></a>

#### `namespace` property

Type: `"string"`. Optional.

Namespace is always system for DNS\_zone.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

<a id="canonical-2202002013200000-1311131113033302-0102333123023011-1213002233020111-3003011303231322-0201020220301132-3002233330130110-1011110212101133"></a>

<a id="canonical-1232300133333020-0323120101230220-2222321310221331-2111323112203120-2223012202121201-2311313310131310-1122020332210313-2233130210130200"></a>

#### `zone_name` property

Type: `"string"`. Optional.

Zone Name. Human-readable name for the resource

<a id="canonical-2321033120023201-3230313223211221-1103313331133121-1010113101103123-3013120233133131-1313101301331231-1212010232323212-0121110302000020"></a>

### All schema paths for `xcsh_dns_zone_delete_cryptokey`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `key_id` | [key_id](actions--dns_zone_delete_cryptokey--reference--group-001.md#canonical-3301211003310303-0111331331031121-0220001010303130-2011230333323112-0231303310301032-2322203123103010-2232201101333130-1232310022130312) |
| `namespace` | [namespace](actions--dns_zone_delete_cryptokey--reference--group-001.md#canonical-3223321031030132-0133021021112212-0112002012020331-3222120003100311-2012020020022101-1023120300332102-2222303003121130-1213333031320003) |
| `zone_name` | [zone_name](actions--dns_zone_delete_cryptokey--reference--group-001.md#canonical-2202002013200000-1311131113033302-0102333123023011-1213002233020111-3003011303231322-0201020220301132-3002233330130110-1011110212101133) |
