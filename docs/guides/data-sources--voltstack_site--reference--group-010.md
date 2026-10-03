---
page_title: "xcsh_voltstack_site reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_voltstack_site reference."
---

# xcsh_voltstack_site reference

<a id="canonical-2213313110020223-0011332022223020-3310200013233110-0012221031220221-1023011112213300-1203030003333122-0330220112113303-3330213112123223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1331032331000210-3311310232022002-1221320122031213-0330302030323332-3220021101113213-1201331032221231-1331221111103333-1302320022330113"></a>

## offline_survivability_mode.no_offline_survivability_mode — no_offline_survivability_mode / 233013302033 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [offline_survivability_mode](data-sources--voltstack_site--reference--group-009.md#canonical-3012201203013310-2202301221032113-3211212303232023-3001230312211120-0223122330211122-3321300212212230-1123030231222122-0310112232022022)
- offline_survivability_mode.no_offline_survivability_mode

<a id="canonical-0120003103332112-2000021230311013-0213331132302302-3211023310223133-1033223000113331-1023332100201230-3031233103003001-1231213222301010"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no offline survivability mode.

Upstream description:

This can be used for messages where no values are needed.

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

<a id="canonical-1213023103131202-1120002222101200-1231220132033013-3120111302320233-0022121203313110-2223112302012330-1001003220210302-1012022221122200"></a>

## Direct properties — no_offline_survivability_mode / 233013302033 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1331301223330230-0333033010223022-2020000333320133-2110232101122121-2312030112332103-1323222231222203-3022121320132302-2223023303130203"></a>

## Next pages — no_offline_survivability_mode / 233013302033 / 4

- [offline_survivability_mode](data-sources--voltstack_site--reference--group-009.md#canonical-3012201203013310-2202301221032113-3211212303232023-3001230312211120-0223122330211122-3321300212212230-1123030231222122-0310112232022022)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-3100102213131302-3311133131010110-2212201213202300-0103003312120103-3133113233030313-2213101010303002-2212003321123131-0123323330123023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0030001132211223-0103323133121000-0310021203031213-3211033131232102-1102030202032102-3102002000111230-1220232000132112-1031013021111112"></a>

## os — os / 022013321302 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- os

<a id="canonical-2202230300313322-2212210010020112-3211100130102130-3033323101110121-0330103202002103-2302000132021012-3223321020332230-0133133301003110"></a>

Type: `"single"`. Computed.

Select the F5XC Operating System Version for the site. By default, latest available OS Version will
be used. Refer to release notes to find required released OS versions.

Upstream description:

Select the F5XC Operating System Version for the site. By default, latest available OS Version will
be used. Refer to release notes to find required released OS versions.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-operating_system_version_choice": "[\"default_os_version\",\"operating_system_version\"]"
}
```

<a id="canonical-1011233323122331-2120132001220323-0001322231123120-2031123232113203-2111120311122333-1302110301131110-3221212302021031-3120023022012123"></a>

## Direct properties — os / 022013321302 / 3

- [default_os_version](data-sources--voltstack_site--reference--group-010.md#canonical-3330213213300202-2200233323023130-2310230320230102-2032320232121202-2303011321102223-1000203001100223-1131032133020133-1232312111102121): complete subsection reference.

<a id="canonical-0030332133000102-0211211031222333-1130012121203232-2332303112200130-0101211330300013-2232222002211311-0221221003231323-0203220030313330"></a>

<a id="canonical-1212311003103223-2033102221231133-1100002100302333-3000303322303210-1212213301233100-0302233022101320-3302003332122211-1312112031323030"></a>

## operating_system_version property — os / 022013321302 / 4

Type: `"string"`. Computed.

Exclusive with \[default\_os\_version\] Specify a OS version to be used e.g. 9.2024.6.

Upstream description:

Exclusive with \[default\_os\_version\] Specify a OS version to be used e.g. 9.2024.6.

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
    "ves.io.schema.rules.string.max_len": "20"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "20"
  }
}
```

<a id="canonical-3312123012132223-3322330110031110-2130001111132310-2222303302203333-0033320231212233-0032002131200322-0321120312001201-0013001131223031"></a>

## Next pages — os / 022013321302 / 5

- [os.default_os_version](data-sources--voltstack_site--reference--group-010.md#canonical-3330213213300202-2200233323023130-2310230320230102-2032320232121202-2303011321102223-1000203001100223-1131032133020133-1232312111102121)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-3330213213300202-2200233323023130-2310230320230102-2032320232121202-2303011321102223-1000203001100223-1131032133020133-1232312111102121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1121223312303122-1020213222110033-0330010330233322-0301133023113333-2301020122230220-2330103011310103-0120221233213111-0012120003100301"></a>

## os.default_os_version — default_os_version / 030012111130 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [os](data-sources--voltstack_site--reference--group-010.md#canonical-3100102213131302-3311133131010110-2212201213202300-0103003312120103-3133113233030313-2213101010303002-2212003321123131-0123323330123023)
- os.default_os_version

<a id="canonical-3322000231000230-0322213002003122-0323110111022112-0231213313301131-1030100222022130-3102223003121002-2212030321103202-2201333331033111"></a>

Type: `["object", {}]`. Computed.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

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

<a id="canonical-0000231210102130-2022100031333233-0011020312132012-2311310000101203-1000303030101233-0013203202201223-1200201033222111-2021130002131103"></a>

## Direct properties — default_os_version / 030012111130 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3130200032031222-1321012230301010-2211131201120123-0023031201333133-1132012112203311-2011333331331133-3021120203031121-0223331313131313"></a>

## Next pages — default_os_version / 030012111130 / 4

- [os](data-sources--voltstack_site--reference--group-010.md#canonical-3100102213131302-3311133131010110-2212201213202300-0103003312120103-3133113233030313-2213101010303002-2212003321123131-0123323330123023)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-1103213303213211-0311100031033111-3302230130111311-2303232333310320-0133310202311332-2311222201011323-3120221112120033-0011312001223100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1211130212300123-0100130112300111-1312210301133301-2230303331212102-3022111122120122-2223110331033320-0301123133111103-2230102221202210"></a>

## sriov_interfaces — sriov_interfaces / 111213033113 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- sriov_interfaces

<a id="canonical-2130201231213303-1230003033112020-0213313101120303-2003323211110112-1013013211313330-2012210030303101-3100322023220312-0232331302001023"></a>

Type: `"single"`. Computed.

List of all custom SR-IOV interfaces configuration.

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

<a id="canonical-1123001032322111-0333331022321021-0212111132301313-0211221112031211-3130021130021232-1000232021122032-0213231130231010-0231031113312132"></a>

## Direct properties — sriov_interfaces / 111213033113 / 3

- [sriov_interface](data-sources--voltstack_site--reference--group-010.md#canonical-1120121332001033-0313310122003230-2033113330131331-2122201132021132-1012131332111232-1233300013012020-2312333230130011-1113311332211223): complete subsection reference.

<a id="canonical-2201211312112133-3212030210110033-1301312131203012-3130113302110232-1101311231211213-3313001330300123-3330030210200302-2111202130301201"></a>

## Next pages — sriov_interfaces / 111213033113 / 4

- [sriov_interfaces.sriov_interface](data-sources--voltstack_site--reference--group-010.md#canonical-1120121332001033-0313310122003230-2033113330131331-2122201132021132-1012131332111232-1233300013012020-2312333230130011-1113311332211223)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-1120121332001033-0313310122003230-2033113330131331-2122201132021132-1012131332111232-1233300013012020-2312333230130011-1113311332211223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3323220331003223-0301301103100220-3203321112332200-2131213332120102-2212021112022210-3110221012122313-0233122031300221-0022002033023210"></a>

## sriov_interfaces.sriov_interface — sriov_interface / 031320120313 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [sriov_interfaces](data-sources--voltstack_site--reference--group-010.md#canonical-1103213303213211-0311100031033111-3302230130111311-2303232333310320-0133310202311332-2311222201011323-3120221112120033-0011312001223100)
- sriov_interfaces.sriov_interface

<a id="canonical-3022210212120101-3231332033032102-0203210111303202-2221210212121123-2203122331002030-1333000032022103-2223130303212131-2133122000202133"></a>

Type: `"list"`. Computed.

Use custom SR-IOV interfaces Configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3102132030222322-2321121021002130-1322320001012013-3121011332110322-0221133123102030-1003120023133000-3013121113301120-3030231003021221"></a>

## Direct properties — sriov_interface / 031320120313 / 3

<a id="canonical-3012001312212103-1012322312030032-2103230200122131-2110021202301202-2123021030232021-1332131023121010-1033303320302301-0022331131120033"></a>

<a id="canonical-1331222132020330-2101111020230133-3002302232311220-1212031103010031-1100122023000100-3203321112123320-2130233012233101-0113013120030320"></a>

## interface_name property — sriov_interface / 031320120313 / 4

Type: `"string"`. Computed.

Name of physical interface. Name of SR-IOV physical interface.

Upstream description:

Name of SR-IOV physical interface.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-2210311100300033-3333121322202132-2133312203230102-2020030030201113-0031322201003331-1011132120201221-3031203313323031-3030302221320220"></a>

<a id="canonical-1321032001303303-2002203133032121-1213002111303013-3010130111210333-3332233300023303-2130013302310100-2021320202131231-0320332311331032"></a>

## number_of_vfio_vfs property — sriov_interface / 031320120313 / 5

Type: `"number"`. Computed.

Number of virtual functions reserved for VNFs and DPDK-based CNFs.

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

<a id="canonical-2333121122101200-0000320121211331-2120323100211122-2320220222001312-1121100132013220-2312223133203122-3200001212220312-3213311103233013"></a>

<a id="canonical-0320022112230102-3311030312312310-1122033220320001-1321120102233330-3210103123131233-1132330201200313-2030022212231321-2222110322022032"></a>

## number_of_vfs property — sriov_interface / 031320120313 / 6

Type: `"number"`. Computed.

Total number of virtual functions. Total number of virtual functions.

Upstream description:

Total number of virtual functions.

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
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-2021312221233022-0310311030022101-0102122013310003-0023121001030133-2223202311022323-0100130310203122-2310233202312022-2103011313011013"></a>

## Next pages — sriov_interface / 031320120313 / 7

- [sriov_interfaces](data-sources--voltstack_site--reference--group-010.md#canonical-1103213303213211-0311100031033111-3302230130111311-2303232333310320-0133310202311332-2311222201011323-3120221112120033-0011312001223100)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-3133230301232221-1223023111221301-1202112211133303-2202023113310010-3011203311032120-0303120202211331-2113302331010032-2112301303323312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3221022201333212-1111030113002123-1023000032010332-0323300332122102-3103110032303303-1013133012200031-1010011031120122-2011213010203310"></a>

## sw — sw / 333010121013 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- sw

<a id="canonical-2131230230123130-3231023002002122-1013022000213300-0003233021310222-2333113013120101-1021030113030022-1322103000022130-0120011332103233"></a>

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

<a id="canonical-3301212323302120-3230110233103323-1120320303011020-3233011103321200-3031011323122201-2231211010103033-2120331100203102-0123021212032131"></a>

## Direct properties — sw / 333010121013 / 3

- [default_sw_version](data-sources--voltstack_site--reference--group-010.md#canonical-3323000030312033-1300300031303110-0300222132011332-2100203122330112-2131103000123030-3300203302300013-1003121311023333-2213331331202120): complete subsection reference.

<a id="canonical-1212220301200121-3020321222221130-2300333232310110-0002302003301112-1000312201213211-2201321120320020-3200332332003302-0330203032021132"></a>

<a id="canonical-0002231112103001-1303001032233330-1223223021321310-1330323100023211-3113103300333201-1103110013113213-2230311222323333-1120310331013221"></a>

## volterra_software_version property — sw / 333010121013 / 4

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
    "ves.io.schema.rules.string.max_len": "20"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "20"
  }
}
```

<a id="canonical-0023330213020021-1010023112012030-3332221121102011-3123321102020312-1031231031202120-1213132321212211-3310013000321303-0322212003002123"></a>

## Next pages — sw / 333010121013 / 5

- [sw.default_sw_version](data-sources--voltstack_site--reference--group-010.md#canonical-3323000030312033-1300300031303110-0300222132011332-2100203122330112-2131103000123030-3300203302300013-1003121311023333-2213331331202120)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-3323000030312033-1300300031303110-0300222132011332-2100203122330112-2131103000123030-3300203302300013-1003121311023333-2213331331202120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2022222010232023-1123122131000132-2010012003113031-2010300030203300-0202130201100323-2121001103301310-3312312103311020-3032101223023113"></a>

## sw.default_sw_version — default_sw_version / 120022020102 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [sw](data-sources--voltstack_site--reference--group-010.md#canonical-3133230301232221-1223023111221301-1202112211133303-2202023113310010-3011203311032120-0303120202211331-2113302331010032-2112301303323312)
- sw.default_sw_version

<a id="canonical-3330201010220012-3002033232132100-0323210121001132-2310033323300211-1210233332232003-3102202323132020-3323013000322022-3111022230131001"></a>

Type: `["object", {}]`. Computed.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

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

<a id="canonical-2122130020100333-3212200132323021-2321012200302020-3102322032232313-1031023020032331-1002132202313321-1332303323021231-1303322023231302"></a>

## Direct properties — default_sw_version / 120022020102 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2030203300120030-3211233220212200-3210213032321022-2221121031022201-3201010330122331-0300301303031033-3301101112233030-2133133122312323"></a>

## Next pages — default_sw_version / 120022020102 / 4

- [sw](data-sources--voltstack_site--reference--group-010.md#canonical-3133230301232221-1223023111221301-1202112211133303-2202023113310010-3011203311032120-0303120202211331-2113302331010032-2112301303323312)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-1223032330333221-1120211003302212-3132123311030210-2021022313010323-1200101121303220-0221231301121030-0203332322210210-1120030001133101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2311110300311301-0232312103030032-1010233103132112-3013233131122221-1213001223122020-0313011303013013-0030201211130131-1313223030131323"></a>

## usb_policy — usb_policy / 133113010002 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- usb_policy

<a id="canonical-0133310023222321-1223230121033213-2133133200132230-3213322003003222-2323032203033310-0122323102202230-3022332133112002-0330101300101202"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-1203202031022031-1033331200031102-0213033220303022-0311012301023310-3112011203202031-1102112212032101-0333022102030033-3311203312313122"></a>

## Direct properties — usb_policy / 133113010002 / 3

<a id="canonical-0333133231023231-2330330111330303-1002333230113221-3221223212310022-3231232000001233-3332130010300202-2103123203203020-1101131131221210"></a>

<a id="canonical-2010031100103123-2200133021220203-3002212232302212-3232111033103112-0132332302021132-1210231220313323-2130301122032321-3301003010022032"></a>

## name property — usb_policy / 133113010002 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-2001110220302330-0020311000101120-1210210301132333-0130110300132201-3032330030311301-2220232113332203-2312331200301223-3221121220120230"></a>

<a id="canonical-1233323311200101-1032303312020110-3103032023230320-0321301323201010-1100121311330321-0120322220303112-1321031032102212-1311013213011121"></a>

## namespace property — usb_policy / 133113010002 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-0220003312033330-3102000313220001-3210322322033201-3013233201220021-0030323110000121-0221312113123232-0222001101333112-2002100022230303"></a>

<a id="canonical-3212021212030011-1302322120123201-0231120211312122-2131002212013120-2202110323012122-3220100133220030-1033322230233331-3011232032102020"></a>

## tenant property — usb_policy / 133113010002 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-1301101322202322-2300123322112230-0000101013331332-0120213022232001-2012120112200302-3201223123133322-3011120223311001-2310221013133032"></a>

## Next pages — usb_policy / 133113010002 / 7

- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-1131011312303133-1023302002122003-1333133200020121-3021223030101331-1222202302102321-1033021231101223-0223110202000230-0231001012033213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1203200133102010-0132231013201331-0010131113011021-2321300311313013-3012002321301302-2200022233330102-2010230300021313-3130001121000301"></a>

## waf_signatures — waf_signatures / 122002010321 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- waf_signatures

<a id="canonical-1212200302221301-3203013132330313-3200312222333310-3330101223001010-1101313013331010-1131320000300033-3013020000311200-1021201211321300"></a>

Type: `"single"`. Computed.

Select F5XC WAF Signatures update mode for the site. By default, new signatures will be applied
manually. Refer to release notes for details about available Signatures update modes.

Upstream description:

Select F5XC WAF Signatures update mode for the site. By default, new signatures will be applied
manually. Refer to release notes for details about available Signatures update modes.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-signatures_update_mode_choice": "[\"automatic\",\"manual\"]"
}
```

<a id="canonical-2332312012233332-1100030121020201-2012201000231111-3131333133220111-0122221130222213-3320313023113300-2133233112130330-2102212323220122"></a>

## Direct properties — waf_signatures / 122002010321 / 3

- [automatic](data-sources--voltstack_site--reference--group-010.md#canonical-0323112230303331-1000212323222220-1000113023010232-1110302310202130-2312201130210030-2032222100210020-0002210003013303-2100103202220013): complete subsection reference.

- [manual](data-sources--voltstack_site--reference--group-010.md#canonical-3310103033311120-2003201010323030-1132023102333200-3110001112330100-3100100110120211-3133113231100130-3021333222121320-0300032032320031): complete subsection reference.

<a id="canonical-0222102302001133-3232112020302012-1213010212030322-0221331031332110-3222332310130012-3102000323023230-2210231313001223-3300103130100213"></a>

## Next pages — waf_signatures / 122002010321 / 4

- [waf_signatures.automatic](data-sources--voltstack_site--reference--group-010.md#canonical-0323112230303331-1000212323222220-1000113023010232-1110302310202130-2312201130210030-2032222100210020-0002210003013303-2100103202220013)
- [waf_signatures.manual](data-sources--voltstack_site--reference--group-010.md#canonical-3310103033311120-2003201010323030-1132023102333200-3110001112330100-3100100110120211-3133113231100130-3021333222121320-0300032032320031)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-0323112230303331-1000212323222220-1000113023010232-1110302310202130-2312201130210030-2032222100210020-0002210003013303-2100103202220013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2301220022020003-1112021103300010-0211033020202213-0011130332200212-1321213030220321-0312000223213100-2001221301113110-0232330302302112"></a>

## waf_signatures.automatic — automatic / 113120003210 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [waf_signatures](data-sources--voltstack_site--reference--group-010.md#canonical-1131011312303133-1023302002122003-1333133200020121-3021223030101331-1222202302102321-1033021231101223-0223110202000230-0231001012033213)
- waf_signatures.automatic

<a id="canonical-2123332030323021-2212230002030300-1131000131020103-2313120010300230-3023301201232210-0031013111123033-3013103311101133-1210100120002013"></a>

Type: `["object", {}]`. Computed.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

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

<a id="canonical-3331032231213311-2200222100122030-1333112331011321-3122233030030011-2032131221310022-3331001203132023-3201330311330100-3130303220031321"></a>

## Direct properties — automatic / 113120003210 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2230003202221120-1020130201212203-2122130300301213-1012313002110023-3310001032300131-1210132012231112-1032023012020001-3122210311103132"></a>

## Next pages — automatic / 113120003210 / 4

- [waf_signatures](data-sources--voltstack_site--reference--group-010.md#canonical-1131011312303133-1023302002122003-1333133200020121-3021223030101331-1222202302102321-1033021231101223-0223110202000230-0231001012033213)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-3310103033311120-2003201010323030-1132023102333200-3110001112330100-3100100110120211-3133113231100130-3021333222121320-0300032032320031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1122231302233003-3313330131131022-1333123223230302-2011001022132330-0330033320322303-1100130210202003-2322222013303130-3031103023322201"></a>

## waf_signatures.manual — manual / 002223330021 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [waf_signatures](data-sources--voltstack_site--reference--group-010.md#canonical-1131011312303133-1023302002122003-1333133200020121-3021223030101331-1222202302102321-1033021231101223-0223110202000230-0231001012033213)
- waf_signatures.manual

<a id="canonical-1202301211230103-3231132300100330-1113133222330022-0332320221312132-3032210213312033-2303212223001022-2323320210020000-3133321300030330"></a>

Type: `["object", {}]`. Computed.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

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

<a id="canonical-2131032210110310-3011103230011123-3120310112022333-2313021000033102-0032003310311320-0130230310131302-0233103323302302-3012020333133031"></a>

## Direct properties — manual / 002223330021 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2302012020021020-0130001133120033-1223031002223220-3331113100322030-2003013210131122-0001320100201000-2110011220120002-1032222301103131"></a>

## Next pages — manual / 002223330021 / 4

- [waf_signatures](data-sources--voltstack_site--reference--group-010.md#canonical-1131011312303133-1023302002122003-1333133200020121-3021223030101331-1222202302102321-1033021231101223-0223110202000230-0231001012033213)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
