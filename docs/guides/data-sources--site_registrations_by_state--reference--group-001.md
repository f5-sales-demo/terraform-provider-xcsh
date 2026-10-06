---
page_title: "xcsh_site_registrations_by_state reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_site_registrations_by_state reference."
---

# xcsh_site_registrations_by_state reference

<a id="canonical-0320011213231301-1102033233211220-0302131120233130-2122121213220203-2033211132132211-0333231313221113-3210300123023021-2021122221230130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_site_registrations_by_state](../data-sources/site_registrations_by_state.md#canonical-1011310232230021-3303123312003111-3120023223111201-0320322111010320-3032200132321200-3310101021222131-0120032131300331-0201133223313303)
- Property reference

<a id="canonical-3210303323111122-2230321200010101-2200030322200031-2233131331213021-2303130013333101-2321111231120200-2303332330230030-3311221213223300"></a>

### Direct properties for `xcsh_site_registrations_by_state`

- [errors](data-sources--site_registrations_by_state--reference--group-001.md#canonical-3123201032113231-1110203102030102-0321102211032320-2011122333202121-2011102030213113-2300113212013233-3311310101312121-2030132221300302): complete subsection reference.

- [items](data-sources--site_registrations_by_state--reference--group-001.md#canonical-1133122031213013-2312233003121321-2003030111222220-0123212333122102-1113033120112122-1330232222330233-1111303223103223-2020000130121230): complete subsection reference.

<a id="canonical-1311001021113033-0013020103210231-0132213023213212-0331132313231220-3211102013021322-0023032320202320-2231220311230021-3332220303110233"></a>

<a id="canonical-2220301211322233-3030331310222121-2310303232033023-0220330223321133-3330321021131301-0003022213030211-1000102032210010-1330332323232112"></a>

#### `namespace` property

Type: `"string"`. Optional, Computed.

Namespace. Registration namespace, only 'system' namespaces is accepted.

<a id="canonical-3131223012333110-0130012003031211-1120300301022230-1000230123130111-2320100001031112-0231313202200313-3123133123000213-0000330130133310"></a>

<a id="canonical-3212113110330130-1302323021122103-3221311212301031-3113113131022301-1332311132023202-1311120230221303-1131000210212223-3012333100233130"></a>

#### `state` property

Type: `"string"`. Required.

\[Enum:
NOTSET|NEW|APPROVED|ADMITTED|RETIRED|FAILED|DONE|PENDING|ONLINE|UPGRADING|MAINTENANCE|FAILED\_INACTIVE\]
Defines states for registration object State isn't set Object was created (registration request was
received and object created) Registration was approved and waiting for configuration This state can
be set by user only if current state is NEW Registration is approved and prepared for to connect..
Possible values are \`NOTSET\`, \`NEW\`, \`APPROVED\`, \`ADMITTED\`, \`RETIRED\`, \`FAILED\`,
\`DONE\`, \`PENDING\`, \`ONLINE\`, \`UPGRADING\`, \`MAINTENANCE\`, \`FAILED\_INACTIVE\`. Defaults to
\`NOTSET\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["ADMITTED","APPROVED","DONE","FAILED","FAILED_INACTIVE","MAINTENANCE","NEW","NOTSET","ONLINE","PENDING","RETIRED","UPGRADING"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("NOTSET",
    "NEW",
    "APPROVED",
    "ADMITTED",
    "RETIRED",
    "FAILED",
    "DONE",
    "PENDING",
    "ONLINE",
    "UPGRADING",
    "MAINTENANCE",
    "FAILED_INACTIVE"),
}
```

<a id="canonical-3320332322230213-2303331033011103-0110120230302331-3321121332231311-3003311122022030-2123110110211113-1033031001211301-2023101121302000"></a>

### All schema paths for `xcsh_site_registrations_by_state`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `errors` | [errors](data-sources--site_registrations_by_state--reference--group-001.md#canonical-0233103032201323-0133210100320323-0123313010223300-2222021111020023-0103110100010333-0323130331130120-3020300121132130-2332200221113102) |
| `errors.code` | [errors.code](data-sources--site_registrations_by_state--reference--group-001.md#canonical-0211303333302112-2032313300321130-1233002311031101-2132030133232303-0212310130021323-0322222132210313-2211310033130121-1002132030020322) |
| `errors.error_obj` | [errors.error_obj](data-sources--site_registrations_by_state--reference--group-001.md#canonical-2022310132031021-3031031023021221-0113302330030003-3100223213020000-2201312000111122-0021320020103322-3032311302020102-3321301300120220) |
| `errors.error_obj.type_url` | [errors.error_obj.type_url](data-sources--site_registrations_by_state--reference--group-001.md#canonical-3010100033003310-1133011021213222-2331210301103300-1333232130123233-1200020132302333-2221132133223330-0002221120130303-1032310310333031) |
| `errors.error_obj.value` | [errors.error_obj.value](data-sources--site_registrations_by_state--reference--group-001.md#canonical-1000022001002300-0003223321311221-3320312021101131-2031022232310101-1030210320100023-0033222210203033-3231311220213332-1220022010331133) |
| `errors.message` | [errors.message](data-sources--site_registrations_by_state--reference--group-001.md#canonical-0102232032001010-1330123103321012-0020230300023010-2311330130311301-3221201331021212-0101323331313213-1021311321330002-2212030300022210) |
| `items` | [items](data-sources--site_registrations_by_state--reference--group-001.md#canonical-2323020000030220-2023122022323311-1120031030030330-1311031111231111-3102210202310313-1111311200122133-3103213332021312-0033132021100303) |
| `items.annotations` | [items.annotations](data-sources--site_registrations_by_state--reference--group-001.md#canonical-0332101012313233-0300112001301102-3313112333203302-3121131130000223-2232202213211012-2132101110212010-0122232103331120-2220132020011100) |
| `items.description_spec` | [items.description_spec](data-sources--site_registrations_by_state--reference--group-001.md#canonical-3233022003233203-3022033321311213-1323020213130021-0021212211033020-2331010022330303-3202232232031003-3320022221110200-3122022230100030) |
| `items.disabled` | [items.disabled](data-sources--site_registrations_by_state--reference--group-001.md#canonical-0103322331132121-3233010012332023-1131111122300001-0211111131332133-2011112332130210-3100221002202322-2022210323011132-3100322033010220) |
| `items.get_spec` | [items.get_spec](data-sources--site_registrations_by_state--reference--group-001.md#canonical-2212013122131122-3133132230212122-2001122231031230-2103210323131320-3230030333201001-3001231030012030-2121321103212133-0002002113203330) |
| `items.get_spec.infra` | [items.get_spec.infra](data-sources--site_registrations_by_state--reference--group-001.md#canonical-1002221311011210-1131112301131211-0320102012013000-1213223231322011-0320001222013330-3223110033211232-1322112232130020-1202223333330321) |
| `items.get_spec.infra.availability_zone` | [items.get_spec.infra.availability_zone](data-sources--site_registrations_by_state--reference--group-001.md#canonical-2110001112031023-3001212111001230-2133311133312000-0023132113033323-3031020201203201-0123201101233103-0330000232130332-0200211112300223) |
| `items.get_spec.infra.bond_config` | [items.get_spec.infra.bond_config](data-sources--site_registrations_by_state--reference--group-001.md#canonical-3303001302112112-0231233102030210-1231112300212210-2200030210011223-1331122120203131-2022210013213311-1020233332021000-2333323021311312) |
| `items.get_spec.infra.bond_config.interfaces` | [items.get_spec.infra.bond_config.interfaces](data-sources--site_registrations_by_state--reference--group-001.md#canonical-3220330201100321-1112200101303010-0320313221220132-3303322201111320-0000012110313211-3330230023231301-1302132123202323-2011010130213002) |
| `items.get_spec.infra.bond_config.mode` | [items.get_spec.infra.bond_config.mode](data-sources--site_registrations_by_state--reference--group-001.md#canonical-1311111113310010-0021013003213320-2132131330012123-2303033221221331-0231131003111202-0231003232311001-3003220022100223-2102020331313323) |
| `items.get_spec.infra.bond_config.name` | [items.get_spec.infra.bond_config.name](data-sources--site_registrations_by_state--reference--group-001.md#canonical-1122123313203213-0230032233323111-2120330221002100-2023010103010010-2130133211331120-3210310133331302-0030210113120323-3101310013112021) |
| `items.get_spec.infra.certified_hw` | [items.get_spec.infra.certified_hw](data-sources--site_registrations_by_state--reference--group-001.md#canonical-0111331000230202-1301310131133013-0033132310002300-0032133002031113-3001023200330302-0000102221112222-3121020212200011-0322222132032303) |
| `items.get_spec.infra.domain` | [items.get_spec.infra.domain](data-sources--site_registrations_by_state--reference--group-001.md#canonical-2211131130133310-3102212321321232-2231310210103010-2310311110032213-1010321031222331-1323023133221111-2200333030212000-3000023322220232) |
| `items.get_spec.infra.hostname` | [items.get_spec.infra.hostname](data-sources--site_registrations_by_state--reference--group-001.md#canonical-2320322331101001-1212122110302303-1232330203103010-2332111310331122-2000221010112232-0203230321023202-1333113121113201-3333323012210112) |
| `items.get_spec.infra.hugepages` | [items.get_spec.infra.hugepages](data-sources--site_registrations_by_state--reference--group-001.md#canonical-3032120021203310-1331012033230302-3230030332013003-0023102002322223-3100013303220313-2110012212332232-3010111310201210-1332121023203011) |
| `items.get_spec.infra.hugepages.free` | [items.get_spec.infra.hugepages.free](data-sources--site_registrations_by_state--reference--group-001.md#canonical-2301032013200331-0113201303121231-2101002220010021-2233030223302121-3310100111120110-2300002102211000-3123300002300112-1120011230032102) |
| `items.get_spec.infra.hugepages.page_size` | [items.get_spec.infra.hugepages.page_size](data-sources--site_registrations_by_state--reference--group-001.md#canonical-3110310011232000-3213300011320001-1002212113223031-2120032013311103-0300320122331031-3320003113200111-2002030022103330-0200023212222130) |
| `items.get_spec.infra.hugepages.total` | [items.get_spec.infra.hugepages.total](data-sources--site_registrations_by_state--reference--group-001.md#canonical-2110311003230300-3203311023002010-2303003113211330-3012213131333330-0221032101330033-2132030133000303-0103231033021211-1211120132323312) |
| `items.get_spec.infra.hw_info` | [items.get_spec.infra.hw_info](data-sources--site_registrations_by_state--reference--group-001.md#canonical-0333310021202312-0202220020312133-0002332132033111-2132022013221100-3202322313233011-1230230331231023-3321102113223220-2002021322132131) |
| `items.get_spec.infra.hw_info.bios` | [items.get_spec.infra.hw_info.bios](data-sources--site_registrations_by_state--reference--group-001.md#canonical-0103130310322220-3032113002302323-0002030023120003-3020233013000012-0030001030011331-3301210013031021-1111201323221222-0202101330020313) |
| `items.get_spec.infra.hw_info.bios.date` | [items.get_spec.infra.hw_info.bios.date](data-sources--site_registrations_by_state--reference--group-001.md#canonical-3103113300332023-2122213333000301-3000333211200222-1322212332330022-0031203110202213-2111322221331121-2303001113021210-2013303020110320) |
| `items.get_spec.infra.hw_info.bios.vendor` | [items.get_spec.infra.hw_info.bios.vendor](data-sources--site_registrations_by_state--reference--group-001.md#canonical-1202032230213123-2023202330222201-3100311103022303-0002112211301232-2002111222003110-2023300211102232-0303213333320122-2001331321121003) |
| `items.get_spec.infra.hw_info.bios.version` | [items.get_spec.infra.hw_info.bios.version](data-sources--site_registrations_by_state--reference--group-001.md#canonical-1331201120302122-0011320110112111-3013102313211331-2010202111220112-1000311112311123-1210303122000122-0212320110103013-2020002210133331) |
| `items.get_spec.infra.hw_info.board` | [items.get_spec.infra.hw_info.board](data-sources--site_registrations_by_state--reference--group-001.md#canonical-0213010113120210-3032120313301031-3321220112022032-1212323310131213-0131033002013130-2130102013131130-0000103132020110-3023023022102022) |
| `items.get_spec.infra.hw_info.board.asset_tag` | [items.get_spec.infra.hw_info.board.asset_tag](data-sources--site_registrations_by_state--reference--group-001.md#canonical-2111132221130330-1112211010130323-3221221133003101-0031200112321212-1300221121101302-1033303023223223-0231330012331011-2213011101211321) |
| `items.get_spec.infra.hw_info.board.name` | [items.get_spec.infra.hw_info.board.name](data-sources--site_registrations_by_state--reference--group-001.md#canonical-0032301103220210-0302131300203332-1003023012102023-3102113023101131-3311230022330202-0002013213111320-3220121103102020-2012211212331131) |
| `items.get_spec.infra.hw_info.board.serial` | [items.get_spec.infra.hw_info.board.serial](data-sources--site_registrations_by_state--reference--group-001.md#canonical-1212013112030130-1012112112212321-3102320321232203-2320001323303032-0113133213100322-1333222030311232-0030002023301212-0312232121320331) |
| `items.get_spec.infra.hw_info.board.vendor` | [items.get_spec.infra.hw_info.board.vendor](data-sources--site_registrations_by_state--reference--group-001.md#canonical-0113121021330000-0110301002101330-1213131003020230-0313331201303132-0331323130000010-0212200322303331-0321211230132320-2322120131102030) |
| `items.get_spec.infra.hw_info.board.version` | [items.get_spec.infra.hw_info.board.version](data-sources--site_registrations_by_state--reference--group-001.md#canonical-1220323231200000-3120100132230330-2133003110031212-0013322210133333-1023102030300113-0002120133300023-2231100032311200-0021131111123220) |
| `items.get_spec.infra.hw_info.chassis` | [items.get_spec.infra.hw_info.chassis](data-sources--site_registrations_by_state--reference--group-001.md#canonical-3331100232201021-2300001011220320-1210013021023223-3020312101023133-1032213230222200-3031220200020021-2233231203112110-3213110330213213) |
| `items.get_spec.infra.hw_info.chassis.asset_tag` | [items.get_spec.infra.hw_info.chassis.asset_tag](data-sources--site_registrations_by_state--reference--group-001.md#canonical-2210300013210200-2031013103203131-3312203220210231-0233332030002201-2320310012302211-2121113321322123-0211003233022332-3302223021202332) |
| `items.get_spec.infra.hw_info.chassis.serial` | [items.get_spec.infra.hw_info.chassis.serial](data-sources--site_registrations_by_state--reference--group-001.md#canonical-2132122103200322-1322300230100122-0223332311201222-0133221122221130-1301002221131311-2313310313303012-0213210323122313-0032020032321321) |
| `items.get_spec.infra.hw_info.chassis.type` | [items.get_spec.infra.hw_info.chassis.type](data-sources--site_registrations_by_state--reference--group-001.md#canonical-2321300112102020-0310302231223313-3111303220231002-2013102202332033-1131232233230130-1102111300023012-0310000132210030-1300210013203122) |
| `items.get_spec.infra.hw_info.chassis.vendor` | [items.get_spec.infra.hw_info.chassis.vendor](data-sources--site_registrations_by_state--reference--group-001.md#canonical-1221120330122010-2021130130033131-3323023332330322-1002313210301331-0222223032032013-0201103233022202-3322333123130322-1013210311213003) |
| `items.get_spec.infra.hw_info.chassis.version` | [items.get_spec.infra.hw_info.chassis.version](data-sources--site_registrations_by_state--reference--group-001.md#canonical-3330133000110023-2220001212322331-3313003310033312-1231000131320210-0111123130201032-0211123102011212-1031000000220331-3013021100103300) |
| `items.get_spec.infra.hw_info.cpu` | [items.get_spec.infra.hw_info.cpu](data-sources--site_registrations_by_state--reference--group-001.md#canonical-0310321211201110-1303132132113220-1211322301212103-0311023102300023-0301213211202003-3131131103023131-1213330311023200-3201123311121111) |
| `items.get_spec.infra.hw_info.cpu.cache` | [items.get_spec.infra.hw_info.cpu.cache](data-sources--site_registrations_by_state--reference--group-001.md#canonical-0230022223223110-1130110131323312-2120012033311313-0321330110310323-1132303031202311-0012111133110133-2332201203021231-2102331122333221) |
| `items.get_spec.infra.hw_info.cpu.cores` | [items.get_spec.infra.hw_info.cpu.cores](data-sources--site_registrations_by_state--reference--group-001.md#canonical-2213203213303212-1012003030200213-0131211233032203-0001130123021103-0123303331322221-1303012220333110-0201331130132303-1303330130110333) |
| `items.get_spec.infra.hw_info.cpu.cpus` | [items.get_spec.infra.hw_info.cpu.cpus](data-sources--site_registrations_by_state--reference--group-001.md#canonical-2031032103300221-0002103032033222-3230202001213102-0032002212033010-1112130101313232-2031203000100323-1110323033123322-1223020211311101) |
| `items.get_spec.infra.hw_info.cpu.model` | [items.get_spec.infra.hw_info.cpu.model](data-sources--site_registrations_by_state--reference--group-001.md#canonical-3233323222101212-3131023323202311-0202102030103323-1310003330013333-3230203332000200-0301031130223302-1001131322130220-0233030320202011) |
| `items.get_spec.infra.hw_info.cpu.speed` | [items.get_spec.infra.hw_info.cpu.speed](data-sources--site_registrations_by_state--reference--group-001.md#canonical-1101330301300003-2111102322231312-0032210213322221-2130101013232123-3023111203221230-2301313101313210-0221223202211203-3213121311302032) |
| `items.get_spec.infra.hw_info.cpu.threads` | [items.get_spec.infra.hw_info.cpu.threads](data-sources--site_registrations_by_state--reference--group-001.md#canonical-2302223310310221-0203223022102121-3212031231203212-1302110112303322-1221002032020010-2030103202032312-2303002132001231-3012212212133033) |
| `items.get_spec.infra.hw_info.cpu.vendor` | [items.get_spec.infra.hw_info.cpu.vendor](data-sources--site_registrations_by_state--reference--group-001.md#canonical-1030322002120332-2102131112212311-1130221320202300-0110113303101210-3133221331122021-2032300300230113-3130123013303332-3222300000232330) |
| `items.get_spec.infra.hw_info.gpu` | [items.get_spec.infra.hw_info.gpu](data-sources--site_registrations_by_state--reference--group-001.md#canonical-1000211133003211-1233011031300312-3120201120011003-1001120121231011-1321220023230303-2020222323033220-0100133321320121-1233233312320302) |
| `items.get_spec.infra.hw_info.gpu.cuda_version` | [items.get_spec.infra.hw_info.gpu.cuda_version](data-sources--site_registrations_by_state--reference--group-001.md#canonical-1100120200032110-1313323001123003-3333301232011022-0112123312111023-3303100230011130-3301130023201022-3032312102301201-0221111032013332) |
| `items.get_spec.infra.hw_info.gpu.driver_version` | [items.get_spec.infra.hw_info.gpu.driver_version](data-sources--site_registrations_by_state--reference--group-001.md#canonical-1110110203130132-3103222010111220-3322220323210202-2121001223022102-3101003103231032-3303122112302330-3311113312023133-3223033203333020) |
| `items.get_spec.infra.hw_info.gpu.gpu_device` | [items.get_spec.infra.hw_info.gpu.gpu_device](data-sources--site_registrations_by_state--reference--group-001.md#canonical-3130210131120230-1222222131120103-3232113110102031-3012211323322312-2233231322321122-2211333012333132-2201100130120202-2322312012323101) |
| `items.get_spec.infra.hw_info.gpu.gpu_device.id` | [items.get_spec.infra.hw_info.gpu.gpu_device.id](data-sources--site_registrations_by_state--reference--group-001.md#canonical-3012122111010333-0331132220300313-1322103322103011-2300032212210131-1010212221111031-0002102321233221-0022013120300300-1120222100202221) |
| `items.get_spec.infra.hw_info.gpu.gpu_device.processes` | [items.get_spec.infra.hw_info.gpu.gpu_device.processes](data-sources--site_registrations_by_state--reference--group-001.md#canonical-1232301333120131-1133310211230020-0230223313023310-2313002112332203-3120101022211331-0301023223131211-0223323220320210-2203212010133220) |
| `items.get_spec.infra.hw_info.gpu.gpu_device.product_name` | [items.get_spec.infra.hw_info.gpu.gpu_device.product_name](data-sources--site_registrations_by_state--reference--group-001.md#canonical-0102211020331210-0010203332211003-2011020200220030-0233003012232123-1320113031011103-2103201210012232-0330110321203120-0103021220302221) |
| `items.get_spec.infra.hw_info.kernel` | [items.get_spec.infra.hw_info.kernel](data-sources--site_registrations_by_state--reference--group-001.md#canonical-3022210232033133-3022230012103200-3113331312220033-1131032331030111-1330123330100130-1231030122230020-1200213333221321-0030200233120202) |
| `items.get_spec.infra.hw_info.kernel.architecture` | [items.get_spec.infra.hw_info.kernel.architecture](data-sources--site_registrations_by_state--reference--group-001.md#canonical-2301332113310132-2121313133212023-0101132220102001-2233213220322123-3110203233100130-3131130010012311-1332312200123000-1021130023222033) |
| `items.get_spec.infra.hw_info.kernel.release` | [items.get_spec.infra.hw_info.kernel.release](data-sources--site_registrations_by_state--reference--group-001.md#canonical-1222210020132301-3210130023111201-1301323110202122-3100313112210002-3330010002203032-3021323123122322-1330023022110031-1103120023130200) |
| `items.get_spec.infra.hw_info.kernel.version` | [items.get_spec.infra.hw_info.kernel.version](data-sources--site_registrations_by_state--reference--group-001.md#canonical-2223332112123102-2302323111103133-2103331300301020-0323303001310022-1303010003131321-2310000121200111-1202102333333220-0103110013012332) |
| `items.get_spec.infra.hw_info.memory` | [items.get_spec.infra.hw_info.memory](data-sources--site_registrations_by_state--reference--group-001.md#canonical-0130112022201013-1331113032132201-0003330032103203-3121102022232113-3321331223231213-0313331211101203-1230212112331112-0222200101213223) |
| `items.get_spec.infra.hw_info.memory.size_mb` | [items.get_spec.infra.hw_info.memory.size_mb](data-sources--site_registrations_by_state--reference--group-001.md#canonical-1033131011102131-2213133302000101-3312001301220031-0133113110030120-0212303011103321-0000022030322012-0001233332101202-2032311001100111) |
| `items.get_spec.infra.hw_info.memory.speed` | [items.get_spec.infra.hw_info.memory.speed](data-sources--site_registrations_by_state--reference--group-001.md#canonical-2130110212310321-3220300310113111-2331023030232001-3233103320322111-2311322301210110-2101211120303123-3213203220201130-0201203111322203) |
| `items.get_spec.infra.hw_info.memory.type` | [items.get_spec.infra.hw_info.memory.type](data-sources--site_registrations_by_state--reference--group-001.md#canonical-2312100222112212-3230333033313330-3200002220223301-1120200221121333-1102012101323333-0122003210033003-1210133112121103-1110121231012323) |
| `items.get_spec.infra.hw_info.network` | [items.get_spec.infra.hw_info.network](data-sources--site_registrations_by_state--reference--group-001.md#canonical-3102231021111033-1103120300102101-2212002122332223-2013221212101312-2111222323023200-1323011133301301-1320003132130232-0213331101221223) |
| `items.get_spec.infra.hw_info.network.driver` | [items.get_spec.infra.hw_info.network.driver](data-sources--site_registrations_by_state--reference--group-001.md#canonical-1322021022001323-2023211233222233-2113232223030213-3322231030031322-1211200011010122-1031332022001000-1221200033012323-1321110323022123) |
| `items.get_spec.infra.hw_info.network.ip_address` | [items.get_spec.infra.hw_info.network.ip_address](data-sources--site_registrations_by_state--reference--group-001.md#canonical-2000001002333001-1033002330303333-2233220102122003-2301211300021111-3311011103010102-0113323232321322-2320003302320322-1111112202302231) |
| `items.get_spec.infra.hw_info.network.link_quality` | [items.get_spec.infra.hw_info.network.link_quality](data-sources--site_registrations_by_state--reference--group-001.md#canonical-2302331023020230-0103002311100330-2331203000003101-2120013103302013-1300231032031120-3023011202321031-3220013221331323-2203113102201102) |
| `items.get_spec.infra.hw_info.network.link_type` | [items.get_spec.infra.hw_info.network.link_type](data-sources--site_registrations_by_state--reference--group-001.md#canonical-2301332301233202-0130312323133003-3001021020312003-0211233212023022-1131110103031020-1223321231110210-3010312113132001-0032001022311321) |
| `items.get_spec.infra.hw_info.network.mac_address` | [items.get_spec.infra.hw_info.network.mac_address](data-sources--site_registrations_by_state--reference--group-001.md#canonical-3130032032203332-2330020313100232-2312010130211321-3322131032133023-2021113011322013-1020313303033133-3200030230322103-1130203013133112) |
| `items.get_spec.infra.hw_info.network.name` | [items.get_spec.infra.hw_info.network.name](data-sources--site_registrations_by_state--reference--group-001.md#canonical-1110101333221200-0023201313200201-2302300012031010-0300030201111302-1011030232022302-1220112323231220-2300133002232323-0110132230130110) |
| `items.get_spec.infra.hw_info.network.port` | [items.get_spec.infra.hw_info.network.port](data-sources--site_registrations_by_state--reference--group-001.md#canonical-2110010322100311-1122331232002331-0322332212010121-0222031003000233-0023233030322120-1233131231100302-2201102002302111-0113111313101213) |
| `items.get_spec.infra.hw_info.network.speed` | [items.get_spec.infra.hw_info.network.speed](data-sources--site_registrations_by_state--reference--group-001.md#canonical-2311111111303331-1210003131233313-2232032100231021-3032311032011202-1123012011100211-3312302020230233-3001123103222132-1310020302323203) |
| `items.get_spec.infra.hw_info.numa_nodes` | [items.get_spec.infra.hw_info.numa_nodes](data-sources--site_registrations_by_state--reference--group-001.md#canonical-2103222010122313-2101300233130033-3020113032232000-2122103113221200-2122133223212320-2111001201211110-1310113212020330-3022203112301123) |
| `items.get_spec.infra.hw_info.os` | [items.get_spec.infra.hw_info.os](data-sources--site_registrations_by_state--reference--group-001.md#canonical-3130030000102013-2332131113101102-0200331231331111-3112130120211101-3320303202220312-3002320231012122-2013311013022232-3230221001303231) |
| `items.get_spec.infra.hw_info.os.architecture` | [items.get_spec.infra.hw_info.os.architecture](data-sources--site_registrations_by_state--reference--group-001.md#canonical-2211230221112302-3010000203133202-3121202332002233-0303313101323000-0333102022101101-2203303030030123-0120031010110022-2022000111310312) |
| `items.get_spec.infra.hw_info.os.name` | [items.get_spec.infra.hw_info.os.name](data-sources--site_registrations_by_state--reference--group-001.md#canonical-1110231003230230-2212313002303202-1313012030133023-0320311333132333-3010102211123121-1132010222012302-0200111002013003-3003021022330231) |
| `items.get_spec.infra.hw_info.os.release` | [items.get_spec.infra.hw_info.os.release](data-sources--site_registrations_by_state--reference--group-001.md#canonical-1002103113013123-2103101322031010-3330211102320232-0200301200222120-2030211332003233-0010312302022123-0022213221130032-3110323033132102) |
| `items.get_spec.infra.hw_info.os.vendor` | [items.get_spec.infra.hw_info.os.vendor](data-sources--site_registrations_by_state--reference--group-001.md#canonical-0223310213023122-1300132030302221-0000020333233000-1200100222211022-3322003012210222-1000333122330103-3233233210220023-1002002330101331) |
| `items.get_spec.infra.hw_info.os.version` | [items.get_spec.infra.hw_info.os.version](data-sources--site_registrations_by_state--reference--group-001.md#canonical-2322312213113031-2033103210201303-1333210333320010-2003121321323102-2111122131120033-0121320233122223-2201000120330010-1122013320030130) |
| `items.get_spec.infra.hw_info.product` | [items.get_spec.infra.hw_info.product](data-sources--site_registrations_by_state--reference--group-001.md#canonical-1223313123203113-3320111012322010-2302013000111230-3211022213332132-2333123033123122-2012020221310313-1021021323232202-1221110021131200) |
| `items.get_spec.infra.hw_info.product.name` | [items.get_spec.infra.hw_info.product.name](data-sources--site_registrations_by_state--reference--group-001.md#canonical-0312113101230321-1020300101000111-3110332232222030-2112311232320011-1303212120330102-1101023122233322-3201231202112023-1300333213102213) |
| `items.get_spec.infra.hw_info.product.serial` | [items.get_spec.infra.hw_info.product.serial](data-sources--site_registrations_by_state--reference--group-001.md#canonical-3130330123221303-3211020003012013-1012132332030313-1322203233333201-0322103023130111-3001222120330001-1320211332312200-1200302121212020) |
| `items.get_spec.infra.hw_info.product.vendor` | [items.get_spec.infra.hw_info.product.vendor](data-sources--site_registrations_by_state--reference--group-001.md#canonical-2121132001111222-1222121021103311-3233020022102332-2333231331100213-0212112232030322-1222001330330330-1203011122311102-3003320331021032) |
| `items.get_spec.infra.hw_info.product.version` | [items.get_spec.infra.hw_info.product.version](data-sources--site_registrations_by_state--reference--group-001.md#canonical-1020213222201123-3322323121110223-0122010302203201-3313131311301100-3123010331312330-0202010113131300-2233311312312312-0010001020202310) |
| `items.get_spec.infra.hw_info.storage` | [items.get_spec.infra.hw_info.storage](data-sources--site_registrations_by_state--reference--group-002.md#canonical-0120003111002213-1232010213213301-0022313003133112-0102031001312033-1133020300023110-2210201030011100-0002203120000001-2202011211000302) |
| `items.get_spec.infra.hw_info.storage.driver` | [items.get_spec.infra.hw_info.storage.driver](data-sources--site_registrations_by_state--reference--group-002.md#canonical-3103230203123023-3110020333222102-1332013010231331-0213032020100232-2223210133201313-2233030002003133-3332301012203303-1030200203102310) |
| `items.get_spec.infra.hw_info.storage.model` | [items.get_spec.infra.hw_info.storage.model](data-sources--site_registrations_by_state--reference--group-002.md#canonical-1032233102023333-1001022120331330-1131002110032222-2310333312000113-3311010030102010-3112100212003132-0100213031012012-3033132021210023) |
| `items.get_spec.infra.hw_info.storage.name` | [items.get_spec.infra.hw_info.storage.name](data-sources--site_registrations_by_state--reference--group-002.md#canonical-0123121022213001-3122002330300011-2233211031121000-2301021133131123-1133013332103001-1020001121120033-0103033123132200-3223110121231101) |
| `items.get_spec.infra.hw_info.storage.serial` | [items.get_spec.infra.hw_info.storage.serial](data-sources--site_registrations_by_state--reference--group-002.md#canonical-2112321113001220-0121220110201322-3132002121112020-0022131030001023-3102102231002000-1212332302003130-0323221013002303-2332303300011002) |
| `items.get_spec.infra.hw_info.storage.size_gb` | [items.get_spec.infra.hw_info.storage.size_gb](data-sources--site_registrations_by_state--reference--group-002.md#canonical-0102230202111133-1322313003233021-1133322210032310-2202232113211123-1123121331201323-3212321333022301-1123122211210313-3320223020011321) |
| `items.get_spec.infra.hw_info.storage.vendor` | [items.get_spec.infra.hw_info.storage.vendor](data-sources--site_registrations_by_state--reference--group-002.md#canonical-1331303310200111-0011032121020210-2010023021203131-0030222300302121-0201111112130333-0323112221202101-2030120313323303-2011221301211220) |
| `items.get_spec.infra.hw_info.usb` | [items.get_spec.infra.hw_info.usb](data-sources--site_registrations_by_state--reference--group-002.md#canonical-2000030012113221-3003101022101230-3203111132133312-2102300131212023-0311022110212022-2202301033020200-3002222133221121-2020301320230220) |
| `items.get_spec.infra.hw_info.usb.address` | [items.get_spec.infra.hw_info.usb.address](data-sources--site_registrations_by_state--reference--group-002.md#canonical-3231313113101311-2303211101332202-2231023022220122-0201322203310301-2212211121220113-3200120322023301-3210023202220203-2200221010121322) |
| `items.get_spec.infra.hw_info.usb.b_device_class` | [items.get_spec.infra.hw_info.usb.b_device_class](data-sources--site_registrations_by_state--reference--group-002.md#canonical-0000020101023030-1213301311003213-0112001202123120-0323230133133110-1121131023021131-3321002110303133-2231002131300213-2220031211313331) |
| `items.get_spec.infra.hw_info.usb.b_device_protocol` | [items.get_spec.infra.hw_info.usb.b_device_protocol](data-sources--site_registrations_by_state--reference--group-002.md#canonical-3013213001002011-3130303102010321-0003331213313303-3123231020301332-1221232302120030-3331012130131212-0003100200030230-1300130132011320) |
| `items.get_spec.infra.hw_info.usb.b_device_sub_class` | [items.get_spec.infra.hw_info.usb.b_device_sub_class](data-sources--site_registrations_by_state--reference--group-002.md#canonical-1301232031211323-0033321203333032-3300332310001103-2122333222101213-1122312331301121-2031113022020223-2332323133202201-1301121023122101) |
| `items.get_spec.infra.hw_info.usb.b_max_packet_size` | [items.get_spec.infra.hw_info.usb.b_max_packet_size](data-sources--site_registrations_by_state--reference--group-002.md#canonical-2333202302320130-0012331322011033-3012131202123333-1101002302120010-3110122023302230-0011101001012013-0001230332020331-1013231101302131) |
| `items.get_spec.infra.hw_info.usb.bcd_device` | [items.get_spec.infra.hw_info.usb.bcd_device](data-sources--site_registrations_by_state--reference--group-002.md#canonical-1233201130013001-0320310233301331-0231311232320303-1023120133312330-0121123201200212-1202233302010310-0301000000103310-1203220213013123) |
| `items.get_spec.infra.hw_info.usb.bcd_usb` | [items.get_spec.infra.hw_info.usb.bcd_usb](data-sources--site_registrations_by_state--reference--group-002.md#canonical-1302033200011323-2202230302110122-0033303320030012-1231031331323111-2300113122001233-2211322003313002-2230103210320210-0103103020001103) |
| `items.get_spec.infra.hw_info.usb.bus` | [items.get_spec.infra.hw_info.usb.bus](data-sources--site_registrations_by_state--reference--group-002.md#canonical-3332213330002030-3330023211101222-2233203221230011-2213230110112110-1213330030221132-3213332202031201-3310332223001311-0302002312311322) |
| `items.get_spec.infra.hw_info.usb.description_spec` | [items.get_spec.infra.hw_info.usb.description_spec](data-sources--site_registrations_by_state--reference--group-002.md#canonical-2213001203013122-2212303013330213-0002303323000021-0023231301303323-3332303101031323-0203131022113012-0122300001322103-2032230232310310) |
| `items.get_spec.infra.hw_info.usb.i_manufacturer` | [items.get_spec.infra.hw_info.usb.i_manufacturer](data-sources--site_registrations_by_state--reference--group-002.md#canonical-1301330133101321-3032012022122303-0122130123033003-2231000123231322-3111113102010112-0233320212331013-0312020100231202-2210111303131303) |
| `items.get_spec.infra.hw_info.usb.i_product` | [items.get_spec.infra.hw_info.usb.i_product](data-sources--site_registrations_by_state--reference--group-002.md#canonical-3303321303322000-2232102322231021-1202212223121132-0333112110111122-0310331031211011-2311021212013303-1313113100203213-3311003312003123) |
| `items.get_spec.infra.hw_info.usb.i_serial` | [items.get_spec.infra.hw_info.usb.i_serial](data-sources--site_registrations_by_state--reference--group-002.md#canonical-3310303003110320-2120322002301011-3311220122020310-3110113021032131-2013210232203000-0002131133132023-1201212212101020-1023200213023213) |
| `items.get_spec.infra.hw_info.usb.id_product` | [items.get_spec.infra.hw_info.usb.id_product](data-sources--site_registrations_by_state--reference--group-002.md#canonical-3031131201231223-3323311213133330-1100311020312202-0030131123223012-3332220012310301-3231230113110002-0212233130022331-1333223100212122) |
| `items.get_spec.infra.hw_info.usb.id_vendor` | [items.get_spec.infra.hw_info.usb.id_vendor](data-sources--site_registrations_by_state--reference--group-002.md#canonical-3300120102110213-3222201313122021-2032303131033332-2010230311331321-3032031002233020-0202200232320022-0313022101202103-3011223223123023) |
| `items.get_spec.infra.hw_info.usb.port` | [items.get_spec.infra.hw_info.usb.port](data-sources--site_registrations_by_state--reference--group-002.md#canonical-1210310003313033-1130332103002032-3230022332231321-0311111332210333-1210222102013313-1002010221210321-1312103232033301-0321333132201322) |
| `items.get_spec.infra.hw_info.usb.product_name` | [items.get_spec.infra.hw_info.usb.product_name](data-sources--site_registrations_by_state--reference--group-002.md#canonical-2131300313303102-2323212212303202-2033322103122332-3310333321132321-2300032103110210-3212020332010022-2330211300021103-3302002113203203) |
| `items.get_spec.infra.hw_info.usb.speed` | [items.get_spec.infra.hw_info.usb.speed](data-sources--site_registrations_by_state--reference--group-002.md#canonical-2313031021123332-2110123322200222-3123232230010202-2211001102101122-3322323120200100-2330203210310332-1032232200113331-0021101331022323) |
| `items.get_spec.infra.hw_info.usb.usb_type` | [items.get_spec.infra.hw_info.usb.usb_type](data-sources--site_registrations_by_state--reference--group-002.md#canonical-0222111111322201-2330310312030111-2130210131032010-2300333131212031-3313310213301231-2133200030120232-0112210123132212-1331021313023132) |
| `items.get_spec.infra.hw_info.usb.vendor_name` | [items.get_spec.infra.hw_info.usb.vendor_name](data-sources--site_registrations_by_state--reference--group-002.md#canonical-2031112012032230-0301331303301002-0121331320021031-0000312133321021-1012333123302313-1023111101031112-1001130300331233-0231101002221132) |
| `items.get_spec.infra.instance_id` | [items.get_spec.infra.instance_id](data-sources--site_registrations_by_state--reference--group-001.md#canonical-1100012210010132-3013000233132301-1220112030020021-0223223211311330-0331033331312103-1103030230312212-3233203012320111-2030223022321113) |
| `items.get_spec.infra.interfaces` | [items.get_spec.infra.interfaces](data-sources--site_registrations_by_state--reference--group-002.md#canonical-2021320121231312-0031002103322002-2212101122302231-1300330322031013-2203032313111230-0010023112121302-2013030131323012-1122010331303031) |
| `items.get_spec.infra.internet_proxy` | [items.get_spec.infra.internet_proxy](data-sources--site_registrations_by_state--reference--group-002.md#canonical-0230201030010023-1301101010132321-1211101220121002-3221133333101200-3232222022311230-1333113202223132-0220330100330313-1311113301001301) |
| `items.get_spec.infra.internet_proxy.http_proxy` | [items.get_spec.infra.internet_proxy.http_proxy](data-sources--site_registrations_by_state--reference--group-002.md#canonical-2200231231320131-1013103001232212-0100011232222013-3233133123102223-3333031032013222-2302101020212223-0120330031121110-0012322002023210) |
| `items.get_spec.infra.internet_proxy.https_proxy` | [items.get_spec.infra.internet_proxy.https_proxy](data-sources--site_registrations_by_state--reference--group-002.md#canonical-3002010130103333-1330331123301313-1313123122332033-3130323232232121-0022323100013003-1133222021002032-0212120310230122-0323331303301102) |
| `items.get_spec.infra.internet_proxy.no_proxy` | [items.get_spec.infra.internet_proxy.no_proxy](data-sources--site_registrations_by_state--reference--group-002.md#canonical-3310220322223112-3112203223033031-3112111223102132-2102022022321032-1212311313103202-2301323201001130-3233211033011320-1011212110021320) |
| `items.get_spec.infra.internet_proxy.proxy_cacert_url` | [items.get_spec.infra.internet_proxy.proxy_cacert_url](data-sources--site_registrations_by_state--reference--group-002.md#canonical-2231301130201011-2131332303301233-1210000112233220-0320030210231112-3131100110020101-0231311222131223-0232323201202010-0302200223332330) |
| `items.get_spec.infra.is_slo_static` | [items.get_spec.infra.is_slo_static](data-sources--site_registrations_by_state--reference--group-001.md#canonical-0221312133302033-0312322330000112-2212223133320221-3331231311012001-1131310113001113-2022232313231013-1311111332231201-1223331223130101) |
| `items.get_spec.infra.machine_id` | [items.get_spec.infra.machine_id](data-sources--site_registrations_by_state--reference--group-001.md#canonical-2312101011230113-3012011001232032-2202203322201111-3311210032003121-3332000020033111-0312111213020221-1003212102132322-1323233013003003) |
| `items.get_spec.infra.provider_ref` | [items.get_spec.infra.provider_ref](data-sources--site_registrations_by_state--reference--group-001.md#canonical-0211022311130320-1332130001013222-0013030100303302-1222303121212333-3022323323210110-1201020002133323-0001103302211210-3232221133102123) |
| `items.get_spec.infra.sw_info` | [items.get_spec.infra.sw_info](data-sources--site_registrations_by_state--reference--group-002.md#canonical-1332300323112032-2033201032233130-0023102110301013-0013002123002210-1100211132331320-3213121231301302-1321131030123322-3111323001101032) |
| `items.get_spec.infra.sw_info.sw_version` | [items.get_spec.infra.sw_info.sw_version](data-sources--site_registrations_by_state--reference--group-002.md#canonical-1121030101203113-2110003103010332-2100330100102033-0200222211303011-1023121030203012-0213331022320011-3101312223120231-2313212323233231) |
| `items.get_spec.infra.timestamp` | [items.get_spec.infra.timestamp](data-sources--site_registrations_by_state--reference--group-001.md#canonical-3233110202132322-2230223011110330-2020312101112131-1303000302023133-0301112301100120-1000131112001111-3001200132312223-2002101010202221) |
| `items.get_spec.infra.zone` | [items.get_spec.infra.zone](data-sources--site_registrations_by_state--reference--group-001.md#canonical-3122120203201200-2012113013321233-1022001330220230-3211230102132131-0210223221133031-1011232323200313-0321330030112300-2330300321303033) |
| `items.get_spec.passport` | [items.get_spec.passport](data-sources--site_registrations_by_state--reference--group-002.md#canonical-3011210210300201-0130323311003331-1210122103100003-0221221133120113-2131203313301033-0222223320220002-1223003010202021-0112002213130223) |
| `items.get_spec.passport.cluster_name` | [items.get_spec.passport.cluster_name](data-sources--site_registrations_by_state--reference--group-002.md#canonical-3300300003222130-2212130100313212-3312320333023000-0220033001323301-1131221312003033-3321322330303022-2133113323300122-2123310122311002) |
| `items.get_spec.passport.cluster_size` | [items.get_spec.passport.cluster_size](data-sources--site_registrations_by_state--reference--group-002.md#canonical-3323331302302321-2202202000020003-1030230022200202-2120300300010123-1333333123113011-0312200000020331-1032133313300132-3001320013102300) |
| `items.get_spec.passport.cluster_type` | [items.get_spec.passport.cluster_type](data-sources--site_registrations_by_state--reference--group-002.md#canonical-2022332012012313-1233031022103202-0320233003323100-3302211233033113-2333120232210332-1220021112232011-0022102012220231-2022221202022133) |
| `items.get_spec.passport.default_os_version` | [items.get_spec.passport.default_os_version](data-sources--site_registrations_by_state--reference--group-002.md#canonical-2002330013133120-1230011212030120-1130213210011131-2202300331211003-3021111322013110-0011130100303300-1333211322312321-2100121123130233) |
| `items.get_spec.passport.default_sw_version` | [items.get_spec.passport.default_sw_version](data-sources--site_registrations_by_state--reference--group-002.md#canonical-2201112212312113-3310332223303132-1110230211202120-1220122010103000-2303031230010222-0100330301330230-2320311020130221-0231321313133030) |
| `items.get_spec.passport.latitude` | [items.get_spec.passport.latitude](data-sources--site_registrations_by_state--reference--group-002.md#canonical-2232302202212113-0211133112130100-3232322022101230-1022103110032120-3113021031220102-1331332132023323-3110303130303110-0120003333321320) |
| `items.get_spec.passport.longitude` | [items.get_spec.passport.longitude](data-sources--site_registrations_by_state--reference--group-002.md#canonical-3330221221032012-2200300010232011-2303232220322003-2331200313103330-3311321323021330-2330121310213212-3222323332011100-2322321323031231) |
| `items.get_spec.passport.operating_system_version` | [items.get_spec.passport.operating_system_version](data-sources--site_registrations_by_state--reference--group-002.md#canonical-0330002210301003-0012312111111112-3131032323321112-2103122200231313-1331301232000031-3022310221303320-1303100211222113-3111230113320102) |
| `items.get_spec.passport.private_network_name` | [items.get_spec.passport.private_network_name](data-sources--site_registrations_by_state--reference--group-002.md#canonical-1111230032122030-1230132111210020-1013211300132232-2233333210200201-0203332301322310-2222103311010012-2003010102023212-0202023311002032) |
| `items.get_spec.passport.volterra_software_version` | [items.get_spec.passport.volterra_software_version](data-sources--site_registrations_by_state--reference--group-002.md#canonical-1000201010302100-1123321112331231-0132133101123121-1323131011013111-3103313032332122-0233331023012230-1001300311331302-1102211310010331) |
| `items.get_spec.token` | [items.get_spec.token](data-sources--site_registrations_by_state--reference--group-001.md#canonical-1132231000102130-0300321223021303-0320311233021133-2332111303011102-0300003312323131-3213322113211203-3110021001322030-1201332000301100) |
| `items.labels` | [items.labels](data-sources--site_registrations_by_state--reference--group-002.md#canonical-2122121030030312-0311220300232201-1020032113222222-3330131002031013-2223301013323310-3330303031003021-2000033221002203-0121220203101233) |
| `items.metadata` | [items.metadata](data-sources--site_registrations_by_state--reference--group-002.md#canonical-1223211131323033-0133322301201302-1323303230003032-2311123003122120-1120102333223011-0320133322032221-2321321111303331-1000300301103223) |
| `items.metadata.annotations` | [items.metadata.annotations](data-sources--site_registrations_by_state--reference--group-002.md#canonical-0000311201232103-3232332333230033-2301011222213211-1223121103303330-0300221300031202-1221230303002232-0002033110022021-1320211333202100) |
| `items.metadata.description_spec` | [items.metadata.description_spec](data-sources--site_registrations_by_state--reference--group-002.md#canonical-1332023013010313-2220130331321031-3311133000111311-3323313222123201-1112023333100100-1322102202022321-0311021033010013-0110112022320213) |
| `items.metadata.disable_spec` | [items.metadata.disable_spec](data-sources--site_registrations_by_state--reference--group-002.md#canonical-3002311231023111-1211221202230133-3103233010013311-2300200121011330-3220012020323113-0330020223003312-3111211102012201-1131012203313121) |
| `items.metadata.labels` | [items.metadata.labels](data-sources--site_registrations_by_state--reference--group-002.md#canonical-0023333100101321-1010013103000130-0103321010002322-3233333313212232-1103131201310110-0331013201220222-3221303020311102-2103000200022122) |
| `items.metadata.name` | [items.metadata.name](data-sources--site_registrations_by_state--reference--group-002.md#canonical-3300223331130011-2131222100120133-2111131221111112-1031032231213233-1022201111223303-1310033332000300-0013211200323002-3112323023323101) |
| `items.metadata.namespace` | [items.metadata.namespace](data-sources--site_registrations_by_state--reference--group-002.md#canonical-0112102202201010-2000231030221010-2131110010013113-3103021312113122-1320200101220320-2102230000200300-0221300101011301-3033133103312101) |
| `items.name` | [items.name](data-sources--site_registrations_by_state--reference--group-001.md#canonical-1021003023132032-1011113231020112-0031333132102310-1032320201201131-1213010221023301-0312113013032220-0210313210000120-2203320330312000) |
| `items.namespace` | [items.namespace](data-sources--site_registrations_by_state--reference--group-001.md#canonical-3221000022100322-1203121313020203-1013301011022202-1123332321331130-0230322023303311-0202201202312200-1202231300203333-0213003133130233) |
| `items.object` | [items.object](data-sources--site_registrations_by_state--reference--group-002.md#canonical-3202030001300301-2320111202303313-2313022002300010-1203210331222003-1321010031303312-1110132130223202-0132330013301211-0011132201323012) |
| `items.object.metadata` | [items.object.metadata](data-sources--site_registrations_by_state--reference--group-002.md#canonical-2113022200322202-3130112000220202-1230100102203221-2031101003132131-3211010022222001-1331320323220312-0001002310100212-2232112201200011) |
| `items.object.metadata.annotations` | [items.object.metadata.annotations](data-sources--site_registrations_by_state--reference--group-002.md#canonical-1231013230023000-1133213033021130-0103123113131221-2230102313030013-1221021310100303-1133020210121031-3123220022133031-1112231011212013) |
| `items.object.metadata.description_spec` | [items.object.metadata.description_spec](data-sources--site_registrations_by_state--reference--group-002.md#canonical-3200303332203211-3113231331233112-2332102231121003-1021320221323003-1100203231232200-0301312010023012-1023211030013311-1001212100300023) |
| `items.object.metadata.disable_spec` | [items.object.metadata.disable_spec](data-sources--site_registrations_by_state--reference--group-002.md#canonical-2303113320202133-1021003223000331-1013003102021020-3102211001203212-1130320111212000-1103102133121102-0121311002331210-2011313133233010) |
| `items.object.metadata.labels` | [items.object.metadata.labels](data-sources--site_registrations_by_state--reference--group-002.md#canonical-2223002311212233-0121131013003111-0322211111312221-0333211300323011-3031011102033313-0011321001032012-0112120031231321-3323001320120113) |
| `items.object.metadata.name` | [items.object.metadata.name](data-sources--site_registrations_by_state--reference--group-002.md#canonical-2201102300000133-3230030230011330-2003011103130010-3330002232212302-3203132103110303-0011220033011213-1231303330223000-1322001321030000) |
| `items.object.metadata.namespace` | [items.object.metadata.namespace](data-sources--site_registrations_by_state--reference--group-002.md#canonical-1112330233122210-2332323012213230-1133331131322112-1003221132201301-0203202131023221-0330101332020310-0021012121203212-3320033120000330) |
| `items.object.metadata.uid` | [items.object.metadata.uid](data-sources--site_registrations_by_state--reference--group-002.md#canonical-3130102211331330-0133220123222322-1003132012122030-1113232310303133-0210323001031033-2322123221333111-1223022033123021-2310210103012220) |
| `items.object.spec` | [items.object.spec](data-sources--site_registrations_by_state--reference--group-002.md#canonical-2332213103220130-0000221102030001-0030210211001122-3213112113013220-2030101030121021-2210102032333211-1231232013230130-2213020311022200) |
| `items.object.spec.gc_spec` | [items.object.spec.gc_spec](data-sources--site_registrations_by_state--reference--group-002.md#canonical-3013010110011100-1212021011110302-1103110032020321-0101312012031321-3111121110323100-2030001320320023-2003320121323220-3232312001022003) |
| `items.object.spec.gc_spec.connected_regions` | [items.object.spec.gc_spec.connected_regions](data-sources--site_registrations_by_state--reference--group-002.md#canonical-1110020022032002-1301332103331310-2101032221103011-1123303213332222-0010233013000310-2100122102000210-3011120033302010-1033303113120021) |
| `items.object.spec.gc_spec.infra` | [items.object.spec.gc_spec.infra](data-sources--site_registrations_by_state--reference--group-002.md#canonical-0322012312202012-3203203331212210-2312102122321013-0103123100230000-3100231120023222-1031133030321203-1101331012222101-0122112003120223) |
| `items.object.spec.gc_spec.infra.availability_zone` | [items.object.spec.gc_spec.infra.availability_zone](data-sources--site_registrations_by_state--reference--group-002.md#canonical-1112221031002300-3330133003133130-2230001333103222-1300223320020331-1102300312132111-2132220023132211-1312310300323010-2212113120032302) |
| `items.object.spec.gc_spec.infra.bond_config` | [items.object.spec.gc_spec.infra.bond_config](data-sources--site_registrations_by_state--reference--group-002.md#canonical-2110323200100012-2220001031232122-2123330203133321-2213102320012203-1000031130333322-0131110103211013-3002032102321210-1132113100101313) |
| `items.object.spec.gc_spec.infra.bond_config.interfaces` | [items.object.spec.gc_spec.infra.bond_config.interfaces](data-sources--site_registrations_by_state--reference--group-002.md#canonical-0021223203333013-0311212202011233-2331033333312303-1103212003302323-0333002213323331-1223100231111332-1003110233302120-1322333210030010) |
| `items.object.spec.gc_spec.infra.bond_config.mode` | [items.object.spec.gc_spec.infra.bond_config.mode](data-sources--site_registrations_by_state--reference--group-002.md#canonical-0323212203110211-2320303103320323-3020111330102202-0211111302302223-0331132230322323-3012032323333021-0021200013011100-3213321100100101) |
| `items.object.spec.gc_spec.infra.bond_config.name` | [items.object.spec.gc_spec.infra.bond_config.name](data-sources--site_registrations_by_state--reference--group-002.md#canonical-2211332231123222-1002321203212320-1231222312121320-1122322232223332-0331020210320030-0002203113210330-0202201003122121-1120023132122322) |
| `items.object.spec.gc_spec.infra.certified_hw` | [items.object.spec.gc_spec.infra.certified_hw](data-sources--site_registrations_by_state--reference--group-002.md#canonical-3002230103132001-0210000133103302-0230310101201302-3303002301002232-2002313220233032-3122022201300030-0030233122110013-0212312133022313) |
| `items.object.spec.gc_spec.infra.domain` | [items.object.spec.gc_spec.infra.domain](data-sources--site_registrations_by_state--reference--group-002.md#canonical-1302030302301020-1203210021001303-1222203113303310-2333132311211110-0303211123212131-3210201121222212-1023100313201302-3311103211220031) |
| `items.object.spec.gc_spec.infra.hostname` | [items.object.spec.gc_spec.infra.hostname](data-sources--site_registrations_by_state--reference--group-002.md#canonical-3100103201211010-2233332111112101-2101101032033011-1010123002220323-3122123001023101-3133311322233222-1032230220022121-0320213230111302) |
| `items.object.spec.gc_spec.infra.hugepages` | [items.object.spec.gc_spec.infra.hugepages](data-sources--site_registrations_by_state--reference--group-002.md#canonical-0033301121321111-1101320122022331-2012023023312031-3012332231312320-0000312322333211-3003111332120021-2110210232300230-2301200121312032) |
| `items.object.spec.gc_spec.infra.hugepages.free` | [items.object.spec.gc_spec.infra.hugepages.free](data-sources--site_registrations_by_state--reference--group-002.md#canonical-3300321200021001-2012221322223220-1213010120313000-2203301110321131-1132212030221232-1133112201311120-3103200002133002-1302010033323111) |
| `items.object.spec.gc_spec.infra.hugepages.page_size` | [items.object.spec.gc_spec.infra.hugepages.page_size](data-sources--site_registrations_by_state--reference--group-002.md#canonical-1323103221011120-0001121211011033-2111131013300022-0310111310011320-2223220123330313-2100111123030220-1213302131011210-2320232002013220) |
| `items.object.spec.gc_spec.infra.hugepages.total` | [items.object.spec.gc_spec.infra.hugepages.total](data-sources--site_registrations_by_state--reference--group-002.md#canonical-0223010021113211-1002310112123032-0212222222230333-1003132212230221-0203230202020303-3301220210110322-1130112010120130-3332313301122230) |
| `items.object.spec.gc_spec.infra.hw_info` | [items.object.spec.gc_spec.infra.hw_info](data-sources--site_registrations_by_state--reference--group-002.md#canonical-2320000202232123-1312210002113331-3230101321211032-2203210130112303-3331002322301100-1003301003330010-0002222220311112-0020201033200002) |
| `items.object.spec.gc_spec.infra.hw_info.bios` | [items.object.spec.gc_spec.infra.hw_info.bios](data-sources--site_registrations_by_state--reference--group-002.md#canonical-1331313131110320-0130332322132331-3103102301222323-2002232123301102-0003202101112101-2200220323213112-3100123003222031-2113322101021322) |
| `items.object.spec.gc_spec.infra.hw_info.bios.date` | [items.object.spec.gc_spec.infra.hw_info.bios.date](data-sources--site_registrations_by_state--reference--group-002.md#canonical-1123100030133300-3131121222313223-2231323103101022-3303203310220222-2302202111033323-0121132101032301-0233000313002023-0222220131102103) |
| `items.object.spec.gc_spec.infra.hw_info.bios.vendor` | [items.object.spec.gc_spec.infra.hw_info.bios.vendor](data-sources--site_registrations_by_state--reference--group-002.md#canonical-1313113303000211-2123012321323212-0213232001212320-0331203300200332-2023022310320203-1120311012323101-2022120300231000-3223212321032201) |
| `items.object.spec.gc_spec.infra.hw_info.bios.version` | [items.object.spec.gc_spec.infra.hw_info.bios.version](data-sources--site_registrations_by_state--reference--group-002.md#canonical-0032322322121103-0320310011110133-0131300121102023-3200123113303020-0332310210220300-0310231030322212-3220333310122312-3031330123110321) |
| `items.object.spec.gc_spec.infra.hw_info.board` | [items.object.spec.gc_spec.infra.hw_info.board](data-sources--site_registrations_by_state--reference--group-002.md#canonical-0121011331032313-2323023003011113-2223102302233203-1132013233202212-0002030011012212-3301311123111011-0002130001033332-3303232102002210) |
| `items.object.spec.gc_spec.infra.hw_info.board.asset_tag` | [items.object.spec.gc_spec.infra.hw_info.board.asset_tag](data-sources--site_registrations_by_state--reference--group-002.md#canonical-1333232202211013-1002031013200323-0031211121001023-3100122013203330-2000231210001032-0001232123232030-1310312002213000-2111000013320233) |
| `items.object.spec.gc_spec.infra.hw_info.board.name` | [items.object.spec.gc_spec.infra.hw_info.board.name](data-sources--site_registrations_by_state--reference--group-002.md#canonical-3332212122322230-3120102332100003-1011203220000130-3313112301111213-0002011323233022-0010301123312113-1312001332321110-2221130000012213) |
| `items.object.spec.gc_spec.infra.hw_info.board.serial` | [items.object.spec.gc_spec.infra.hw_info.board.serial](data-sources--site_registrations_by_state--reference--group-002.md#canonical-0303022020023022-2001101310021112-3130123133113303-2122113133012321-1002313111110111-1333200303200121-2201312230103133-0013023320333130) |
| `items.object.spec.gc_spec.infra.hw_info.board.vendor` | [items.object.spec.gc_spec.infra.hw_info.board.vendor](data-sources--site_registrations_by_state--reference--group-002.md#canonical-2303310301302133-2220133210121323-2101231332123203-0210011213222330-3001031222201232-3303031021320002-1310132110112220-1212110111023101) |
| `items.object.spec.gc_spec.infra.hw_info.board.version` | [items.object.spec.gc_spec.infra.hw_info.board.version](data-sources--site_registrations_by_state--reference--group-002.md#canonical-3120111223220333-0131321220102122-3121030000302020-2003132221312032-2200212302020331-0320010321301213-1010211032120201-3230111030211020) |
| `items.object.spec.gc_spec.infra.hw_info.chassis` | [items.object.spec.gc_spec.infra.hw_info.chassis](data-sources--site_registrations_by_state--reference--group-002.md#canonical-1111230032311231-2120021103000210-2021233111320212-1030322321111330-2220330221020232-2001200212130301-1200013331301123-1032112012033220) |
| `items.object.spec.gc_spec.infra.hw_info.chassis.asset_tag` | [items.object.spec.gc_spec.infra.hw_info.chassis.asset_tag](data-sources--site_registrations_by_state--reference--group-002.md#canonical-0020330222221212-3322233222232133-2323122313121311-3203220231121211-1111222013202012-0001203200002110-0331321220232312-0221223113202231) |
| `items.object.spec.gc_spec.infra.hw_info.chassis.serial` | [items.object.spec.gc_spec.infra.hw_info.chassis.serial](data-sources--site_registrations_by_state--reference--group-002.md#canonical-2302033221102322-0130113320211331-3013211031001302-1011013032300333-1211113100331100-0312123312330123-1300210022311322-1021120222211300) |
| `items.object.spec.gc_spec.infra.hw_info.chassis.type` | [items.object.spec.gc_spec.infra.hw_info.chassis.type](data-sources--site_registrations_by_state--reference--group-002.md#canonical-3033031001022301-0120013123321301-1233133303110212-2301213233113003-3323311330203020-0233201101003003-3213130012013113-3110030211321310) |
| `items.object.spec.gc_spec.infra.hw_info.chassis.vendor` | [items.object.spec.gc_spec.infra.hw_info.chassis.vendor](data-sources--site_registrations_by_state--reference--group-002.md#canonical-3322211012023121-1120220101323113-2101002113021233-1002323010002213-3203330010230233-3100111032200310-3123302123203021-1233203132222321) |
| `items.object.spec.gc_spec.infra.hw_info.chassis.version` | [items.object.spec.gc_spec.infra.hw_info.chassis.version](data-sources--site_registrations_by_state--reference--group-002.md#canonical-2231120110100213-2231220110010210-2313012010012122-2121000100103100-0301210221200313-3220131331033032-3221233321223202-3300333031223233) |
| `items.object.spec.gc_spec.infra.hw_info.cpu` | [items.object.spec.gc_spec.infra.hw_info.cpu](data-sources--site_registrations_by_state--reference--group-002.md#canonical-1300112221033122-2023301221310323-3010310301232331-1013010201330030-2023130310102123-3223301031122120-3303320210231010-3313331231220110) |
| `items.object.spec.gc_spec.infra.hw_info.cpu.cache` | [items.object.spec.gc_spec.infra.hw_info.cpu.cache](data-sources--site_registrations_by_state--reference--group-002.md#canonical-0231310301132001-2223220323212300-0022330033123112-2102131221323233-0112311110030100-3003221021220333-1012203203123131-0300333311221222) |
| `items.object.spec.gc_spec.infra.hw_info.cpu.cores` | [items.object.spec.gc_spec.infra.hw_info.cpu.cores](data-sources--site_registrations_by_state--reference--group-002.md#canonical-0003130003302220-3101022213330232-0300100123122223-3212302301222123-0002301201002333-1022101312321122-0210232000222000-2200213210111232) |
| `items.object.spec.gc_spec.infra.hw_info.cpu.cpus` | [items.object.spec.gc_spec.infra.hw_info.cpu.cpus](data-sources--site_registrations_by_state--reference--group-002.md#canonical-3130012322211231-3330120330202232-0201111020010233-1331230303131302-2221201022212132-3002312030222113-1201203310102101-2220032023202023) |
| `items.object.spec.gc_spec.infra.hw_info.cpu.model` | [items.object.spec.gc_spec.infra.hw_info.cpu.model](data-sources--site_registrations_by_state--reference--group-002.md#canonical-0012230010002002-2332003010232102-1130210122102313-0231123130133303-1200320332111233-1021102011232021-0113212120111333-1312022223331302) |
| `items.object.spec.gc_spec.infra.hw_info.cpu.speed` | [items.object.spec.gc_spec.infra.hw_info.cpu.speed](data-sources--site_registrations_by_state--reference--group-002.md#canonical-1231211031332131-2123003001211112-2213333220232312-2011103320022201-3003001323020111-1002102330032013-1131333332002102-2230030211310213) |
| `items.object.spec.gc_spec.infra.hw_info.cpu.threads` | [items.object.spec.gc_spec.infra.hw_info.cpu.threads](data-sources--site_registrations_by_state--reference--group-002.md#canonical-3322300031112312-1013220121003310-2202000302311331-2322311113220211-1023123331232033-1312020212211112-3121012132010331-1131023013112311) |
| `items.object.spec.gc_spec.infra.hw_info.cpu.vendor` | [items.object.spec.gc_spec.infra.hw_info.cpu.vendor](data-sources--site_registrations_by_state--reference--group-002.md#canonical-0203110213333032-2310331010313122-1300313110221303-3013210222032031-2010321003112102-0233300301202221-0312030001230301-1000230030233121) |
| `items.object.spec.gc_spec.infra.hw_info.gpu` | [items.object.spec.gc_spec.infra.hw_info.gpu](data-sources--site_registrations_by_state--reference--group-002.md#canonical-2300321011103201-2111033322313133-3203302212202030-3211030311312330-2200013320010221-2012002022313231-0103122031312231-2122222223010231) |
| `items.object.spec.gc_spec.infra.hw_info.gpu.cuda_version` | [items.object.spec.gc_spec.infra.hw_info.gpu.cuda_version](data-sources--site_registrations_by_state--reference--group-002.md#canonical-0031301111101113-0322010330211200-0111313333003101-1233120321303102-1311022302332133-2121101312110330-3321023023203003-0300021131313011) |
| `items.object.spec.gc_spec.infra.hw_info.gpu.driver_version` | [items.object.spec.gc_spec.infra.hw_info.gpu.driver_version](data-sources--site_registrations_by_state--reference--group-002.md#canonical-2101323311313022-2033000313022310-1302223330302022-3130232103301101-2302031012131102-3112221330201301-0132230000231320-0133023021112131) |
| `items.object.spec.gc_spec.infra.hw_info.gpu.gpu_device` | [items.object.spec.gc_spec.infra.hw_info.gpu.gpu_device](data-sources--site_registrations_by_state--reference--group-002.md#canonical-2111011011002121-3332102130232303-1031322023130221-3220123132331300-2330003331121130-3200313320013233-1231221113021230-1233120003332332) |
| `items.object.spec.gc_spec.infra.hw_info.gpu.gpu_device.id` | [items.object.spec.gc_spec.infra.hw_info.gpu.gpu_device.id](data-sources--site_registrations_by_state--reference--group-002.md#canonical-1112302212121130-3202210101332320-1311122231003321-0223003131122230-3322111112123202-3330121122120210-0001202020332033-2203201200320302) |
| `items.object.spec.gc_spec.infra.hw_info.gpu.gpu_device.processes` | [items.object.spec.gc_spec.infra.hw_info.gpu.gpu_device.processes](data-sources--site_registrations_by_state--reference--group-002.md#canonical-2310102333110032-3023102011122230-3011301310321232-0023033130312032-3032230231221122-0021013313323323-1232203033110102-2233203033223131) |
| `items.object.spec.gc_spec.infra.hw_info.gpu.gpu_device.product_name` | [items.object.spec.gc_spec.infra.hw_info.gpu.gpu_device.product_name](data-sources--site_registrations_by_state--reference--group-002.md#canonical-2231130120332021-1222032321301121-1013120120330312-1001011011303030-2230130003222203-0221331111223203-0121330113212012-2031030033233131) |
| `items.object.spec.gc_spec.infra.hw_info.kernel` | [items.object.spec.gc_spec.infra.hw_info.kernel](data-sources--site_registrations_by_state--reference--group-002.md#canonical-3132133031332030-3032110132033211-0023120320032110-1223201233320321-0201112321203031-1201223310033211-2223221111131000-2312130101221021) |
| `items.object.spec.gc_spec.infra.hw_info.kernel.architecture` | [items.object.spec.gc_spec.infra.hw_info.kernel.architecture](data-sources--site_registrations_by_state--reference--group-002.md#canonical-3121132102032020-0103232330130111-2132323101200132-2002321213313322-1201121011322111-2220321223021100-1103332013302213-1102032300020323) |
| `items.object.spec.gc_spec.infra.hw_info.kernel.release` | [items.object.spec.gc_spec.infra.hw_info.kernel.release](data-sources--site_registrations_by_state--reference--group-002.md#canonical-0212300233331022-2202001203322230-1001123132010022-0231013131013012-2330013210020021-1302001223212001-3332001333203001-1031311210310201) |
| `items.object.spec.gc_spec.infra.hw_info.kernel.version` | [items.object.spec.gc_spec.infra.hw_info.kernel.version](data-sources--site_registrations_by_state--reference--group-002.md#canonical-3203110220232000-2023301033320320-2303322302130201-2223220103001002-1221102330220202-3301113123223132-3111122111030113-1300113333110221) |
| `items.object.spec.gc_spec.infra.hw_info.memory` | [items.object.spec.gc_spec.infra.hw_info.memory](data-sources--site_registrations_by_state--reference--group-002.md#canonical-3300231331130322-1122231220321313-3333001310122122-1322221021331302-2201120213012312-1113033022312130-1122112312100300-0200100030120221) |
| `items.object.spec.gc_spec.infra.hw_info.memory.size_mb` | [items.object.spec.gc_spec.infra.hw_info.memory.size_mb](data-sources--site_registrations_by_state--reference--group-002.md#canonical-1232233100230000-1312032021122231-1111031233132132-0203031220300010-0330301231323223-2212320233303323-1000201101333001-0300322032332030) |
| `items.object.spec.gc_spec.infra.hw_info.memory.speed` | [items.object.spec.gc_spec.infra.hw_info.memory.speed](data-sources--site_registrations_by_state--reference--group-002.md#canonical-2233321100301130-0331223003100112-2330331223130133-2001123222100113-2311121100210132-2112111213110030-3123012231232022-3223310000100110) |
| `items.object.spec.gc_spec.infra.hw_info.memory.type` | [items.object.spec.gc_spec.infra.hw_info.memory.type](data-sources--site_registrations_by_state--reference--group-002.md#canonical-2200022200130310-3222220332300131-0222012332121302-3022102000121321-2213320223022302-0023012220203330-2333002100322033-3223122022011103) |
| `items.object.spec.gc_spec.infra.hw_info.network` | [items.object.spec.gc_spec.infra.hw_info.network](data-sources--site_registrations_by_state--reference--group-002.md#canonical-2113030113310223-2222011102201300-3103230303103103-3333303032221023-1000102103202202-3301202033022222-3132312102012230-0230211113132130) |
| `items.object.spec.gc_spec.infra.hw_info.network.driver` | [items.object.spec.gc_spec.infra.hw_info.network.driver](data-sources--site_registrations_by_state--reference--group-002.md#canonical-3332212131132302-1112003203121132-3203131202303012-1122011032220313-1223332030120033-2332111122113030-1333301332023310-1333101200111002) |
| `items.object.spec.gc_spec.infra.hw_info.network.ip_address` | [items.object.spec.gc_spec.infra.hw_info.network.ip_address](data-sources--site_registrations_by_state--reference--group-002.md#canonical-3110222200313211-2000213020231103-0231330203022323-0032113310233123-2330013312003200-3122233100021001-2002333132220230-0331120133130300) |
| `items.object.spec.gc_spec.infra.hw_info.network.link_quality` | [items.object.spec.gc_spec.infra.hw_info.network.link_quality](data-sources--site_registrations_by_state--reference--group-002.md#canonical-2002202303033301-2332301232211120-2111300030010231-0233031132333300-3033011021313313-2302210332031222-0122323020300331-3111233122312022) |
| `items.object.spec.gc_spec.infra.hw_info.network.link_type` | [items.object.spec.gc_spec.infra.hw_info.network.link_type](data-sources--site_registrations_by_state--reference--group-002.md#canonical-2313312113011200-3320013220101030-2031220132120023-2300111203221220-2232220023300232-1213322320321333-3102323111301001-1102330122120021) |
| `items.object.spec.gc_spec.infra.hw_info.network.mac_address` | [items.object.spec.gc_spec.infra.hw_info.network.mac_address](data-sources--site_registrations_by_state--reference--group-002.md#canonical-1010303121121101-3000320012223203-0033223320312301-1123220330101321-3302031301203022-2202112233000212-2021112122322020-2312100003330222) |
| `items.object.spec.gc_spec.infra.hw_info.network.name` | [items.object.spec.gc_spec.infra.hw_info.network.name](data-sources--site_registrations_by_state--reference--group-002.md#canonical-1211231020123010-3020322333121130-0322230210330322-3312330013310203-2223032001133322-0102012200301200-0313212220112130-0021103303033210) |
| `items.object.spec.gc_spec.infra.hw_info.network.port` | [items.object.spec.gc_spec.infra.hw_info.network.port](data-sources--site_registrations_by_state--reference--group-002.md#canonical-0203301210031013-0021010022002310-1010323231311003-3023313312002231-2210021301321300-1201100213021330-2201122200323200-1122303133220032) |
| `items.object.spec.gc_spec.infra.hw_info.network.speed` | [items.object.spec.gc_spec.infra.hw_info.network.speed](data-sources--site_registrations_by_state--reference--group-003.md#canonical-2123222202103032-0021221112102321-1102333221001111-3133223032333331-3322021211312001-2130313121302100-3312031000101021-2131032012000010) |
| `items.object.spec.gc_spec.infra.hw_info.numa_nodes` | [items.object.spec.gc_spec.infra.hw_info.numa_nodes](data-sources--site_registrations_by_state--reference--group-002.md#canonical-0302203102102231-3222032030230203-2100220323100223-2023303311203232-2133012020121022-1110333100032232-2221303321020212-1002221131111323) |
| `items.object.spec.gc_spec.infra.hw_info.os` | [items.object.spec.gc_spec.infra.hw_info.os](data-sources--site_registrations_by_state--reference--group-003.md#canonical-0131130030221300-1213123123130331-3212332222022302-2022133202101302-3112120122310130-0230303001121012-0132201221120320-3123230333123311) |
| `items.object.spec.gc_spec.infra.hw_info.os.architecture` | [items.object.spec.gc_spec.infra.hw_info.os.architecture](data-sources--site_registrations_by_state--reference--group-003.md#canonical-0223333033033333-2020202203111031-3001103321121122-3003013112123111-3330130210000233-0320210323132303-3013331103123113-2231133213103222) |
| `items.object.spec.gc_spec.infra.hw_info.os.name` | [items.object.spec.gc_spec.infra.hw_info.os.name](data-sources--site_registrations_by_state--reference--group-003.md#canonical-2121221033221311-2311110301233011-0101210110211033-1121000031320120-2332021102330333-0031321000213000-3130212223222130-2212030213301302) |
| `items.object.spec.gc_spec.infra.hw_info.os.release` | [items.object.spec.gc_spec.infra.hw_info.os.release](data-sources--site_registrations_by_state--reference--group-003.md#canonical-0331301331012301-3120123000011221-1300022223121103-2110222102103033-3312200311200022-2010100230310121-0030031033130203-2201331103011222) |
| `items.object.spec.gc_spec.infra.hw_info.os.vendor` | [items.object.spec.gc_spec.infra.hw_info.os.vendor](data-sources--site_registrations_by_state--reference--group-003.md#canonical-2220101323020133-0100133030113013-0132232033122200-3322030310010210-2111010310102121-0302332130303320-1011103013213313-2033002310201110) |
| `items.object.spec.gc_spec.infra.hw_info.os.version` | [items.object.spec.gc_spec.infra.hw_info.os.version](data-sources--site_registrations_by_state--reference--group-003.md#canonical-0030020202302003-2013222312201221-2203300101212002-2101103012032112-3312313100300031-1110301333113113-1210101011231210-0012330011333310) |
| `items.object.spec.gc_spec.infra.hw_info.product` | [items.object.spec.gc_spec.infra.hw_info.product](data-sources--site_registrations_by_state--reference--group-003.md#canonical-3033333103131102-3333113133132303-1211310203023323-0213001031021023-1100111021233313-2100022233010122-0121311321301032-3211123210331011) |
| `items.object.spec.gc_spec.infra.hw_info.product.name` | [items.object.spec.gc_spec.infra.hw_info.product.name](data-sources--site_registrations_by_state--reference--group-003.md#canonical-3022122121101013-0222000003222123-3100212222120002-1312000021001330-2121112311001231-2311122233223030-1320012023202203-3323203132111033) |
| `items.object.spec.gc_spec.infra.hw_info.product.serial` | [items.object.spec.gc_spec.infra.hw_info.product.serial](data-sources--site_registrations_by_state--reference--group-003.md#canonical-1232330030111122-2203011201023123-3111121011133220-3213110332300231-3322321023100122-0133232201223211-1332301110123121-1221002200203200) |
| `items.object.spec.gc_spec.infra.hw_info.product.vendor` | [items.object.spec.gc_spec.infra.hw_info.product.vendor](data-sources--site_registrations_by_state--reference--group-003.md#canonical-2030101313132312-0100311232130313-0222033111103223-0223110310310111-0121233332211231-1223013032002011-3313012330001331-0210122323023022) |
| `items.object.spec.gc_spec.infra.hw_info.product.version` | [items.object.spec.gc_spec.infra.hw_info.product.version](data-sources--site_registrations_by_state--reference--group-003.md#canonical-0120203101301231-2012310122201020-2323231113102013-1221332213221010-2201313332101220-0303230231132012-2020110100213320-1211313221211223) |
| `items.object.spec.gc_spec.infra.hw_info.storage` | [items.object.spec.gc_spec.infra.hw_info.storage](data-sources--site_registrations_by_state--reference--group-003.md#canonical-3001321030133022-3000231102022333-0030312032102033-2210003302113010-0100133330301000-2122231102233231-1310231012203312-1020132222012022) |
| `items.object.spec.gc_spec.infra.hw_info.storage.driver` | [items.object.spec.gc_spec.infra.hw_info.storage.driver](data-sources--site_registrations_by_state--reference--group-003.md#canonical-1003310030133210-2321333133103211-3212202230112200-2122203121001220-2220032210132032-0311111332210020-2111333313101001-1001213111302120) |
| `items.object.spec.gc_spec.infra.hw_info.storage.model` | [items.object.spec.gc_spec.infra.hw_info.storage.model](data-sources--site_registrations_by_state--reference--group-003.md#canonical-1202301303230321-0102201233320111-1000013003030323-3232321112213021-1232133002001120-3333301031132200-2031103130012310-3022220011212001) |
| `items.object.spec.gc_spec.infra.hw_info.storage.name` | [items.object.spec.gc_spec.infra.hw_info.storage.name](data-sources--site_registrations_by_state--reference--group-003.md#canonical-0320333202100131-1312033310330220-0313022031011222-1233000322101021-1202021310110202-2102212221102103-3332020320123002-3133112302222221) |
| `items.object.spec.gc_spec.infra.hw_info.storage.serial` | [items.object.spec.gc_spec.infra.hw_info.storage.serial](data-sources--site_registrations_by_state--reference--group-003.md#canonical-0102020122313130-3322222303122000-1322121013321100-0121021002322212-2320102333303211-0122332011201200-0311231333211013-1212210232132303) |
| `items.object.spec.gc_spec.infra.hw_info.storage.size_gb` | [items.object.spec.gc_spec.infra.hw_info.storage.size_gb](data-sources--site_registrations_by_state--reference--group-003.md#canonical-3331211233003200-1330012213111013-2221121320023313-1233112000221313-1222320033310030-1010301001110112-1012110111210230-1321121322011022) |
| `items.object.spec.gc_spec.infra.hw_info.storage.vendor` | [items.object.spec.gc_spec.infra.hw_info.storage.vendor](data-sources--site_registrations_by_state--reference--group-003.md#canonical-1212100212210313-0330113232132131-2031001101121333-0133122000130221-3111211322103023-2213200202133333-3212221200122103-3031303110123033) |
| `items.object.spec.gc_spec.infra.hw_info.usb` | [items.object.spec.gc_spec.infra.hw_info.usb](data-sources--site_registrations_by_state--reference--group-003.md#canonical-1122000333131123-1300031000000121-3331300122212123-3003112131102011-3132100322022122-1132312122121133-3030210032030023-2033311101121323) |
| `items.object.spec.gc_spec.infra.hw_info.usb.address` | [items.object.spec.gc_spec.infra.hw_info.usb.address](data-sources--site_registrations_by_state--reference--group-003.md#canonical-0330202213030120-3000012131021213-3023301203021211-2023203301230322-2311213121121021-0213301203323132-0133111013103102-2300333012002103) |
| `items.object.spec.gc_spec.infra.hw_info.usb.b_device_class` | [items.object.spec.gc_spec.infra.hw_info.usb.b_device_class](data-sources--site_registrations_by_state--reference--group-003.md#canonical-3022232333233020-0202312023222320-0203102000010010-0103331113121220-1200113001301333-1122212113031123-2212020031003203-1230301313023302) |
| `items.object.spec.gc_spec.infra.hw_info.usb.b_device_protocol` | [items.object.spec.gc_spec.infra.hw_info.usb.b_device_protocol](data-sources--site_registrations_by_state--reference--group-003.md#canonical-2212010303122202-1200212131132322-0023020010100131-2311330310120210-0233202030010131-0123013312123320-1313211012211123-1211202310122131) |
| `items.object.spec.gc_spec.infra.hw_info.usb.b_device_sub_class` | [items.object.spec.gc_spec.infra.hw_info.usb.b_device_sub_class](data-sources--site_registrations_by_state--reference--group-003.md#canonical-3033222232330323-1011123003001032-0011311211012212-1101301231302133-1133113031023220-0023020021203302-3331030221322121-3313023100110310) |
| `items.object.spec.gc_spec.infra.hw_info.usb.b_max_packet_size` | [items.object.spec.gc_spec.infra.hw_info.usb.b_max_packet_size](data-sources--site_registrations_by_state--reference--group-003.md#canonical-3031130303201121-2022320001011032-1021313122330101-1302202003121000-3000020013122113-3111332101002111-2221202031033231-1302112321202032) |
| `items.object.spec.gc_spec.infra.hw_info.usb.bcd_device` | [items.object.spec.gc_spec.infra.hw_info.usb.bcd_device](data-sources--site_registrations_by_state--reference--group-003.md#canonical-1203201330231033-0110202023201331-2022103303131110-0010332101102113-0113311310232300-1112022130022130-0121310331320311-1022222331011203) |
| `items.object.spec.gc_spec.infra.hw_info.usb.bcd_usb` | [items.object.spec.gc_spec.infra.hw_info.usb.bcd_usb](data-sources--site_registrations_by_state--reference--group-003.md#canonical-1322031112010002-0203021132013233-0333020023031202-0021110133311232-0200233010310010-0323330331011103-1030112223002131-3233020231231030) |
| `items.object.spec.gc_spec.infra.hw_info.usb.bus` | [items.object.spec.gc_spec.infra.hw_info.usb.bus](data-sources--site_registrations_by_state--reference--group-003.md#canonical-3210201233232022-0331133313230230-1100322113203130-2331332133301231-1102232311320233-3221013020033112-2221123002333311-2113211100302331) |
| `items.object.spec.gc_spec.infra.hw_info.usb.description_spec` | [items.object.spec.gc_spec.infra.hw_info.usb.description_spec](data-sources--site_registrations_by_state--reference--group-003.md#canonical-0303010303303123-2022211222102013-3301112332002120-3131003010132202-0001322012101301-1233312312020303-3120020031113032-1201310123013012) |
| `items.object.spec.gc_spec.infra.hw_info.usb.i_manufacturer` | [items.object.spec.gc_spec.infra.hw_info.usb.i_manufacturer](data-sources--site_registrations_by_state--reference--group-003.md#canonical-0122231231330203-3002221301000212-1310332012120323-2332312010132221-3230232010320200-3010112222022101-0233332230001113-0221313101202230) |
| `items.object.spec.gc_spec.infra.hw_info.usb.i_product` | [items.object.spec.gc_spec.infra.hw_info.usb.i_product](data-sources--site_registrations_by_state--reference--group-003.md#canonical-3233220022312330-3100100300323221-2303231032201001-0112132001211131-1101330012103201-2010131031211100-1331303123213212-3023010023122313) |
| `items.object.spec.gc_spec.infra.hw_info.usb.i_serial` | [items.object.spec.gc_spec.infra.hw_info.usb.i_serial](data-sources--site_registrations_by_state--reference--group-003.md#canonical-1311130103102002-0200120233311011-0012111213100113-2312330031012231-0011130033203212-0103322312313310-3113302331001322-2332110123111130) |
| `items.object.spec.gc_spec.infra.hw_info.usb.id_product` | [items.object.spec.gc_spec.infra.hw_info.usb.id_product](data-sources--site_registrations_by_state--reference--group-003.md#canonical-1300333022001312-0123322301113101-0023103231101323-3130033031031012-2233123130031333-3021311102110133-1302333212000020-2131121312211013) |
| `items.object.spec.gc_spec.infra.hw_info.usb.id_vendor` | [items.object.spec.gc_spec.infra.hw_info.usb.id_vendor](data-sources--site_registrations_by_state--reference--group-003.md#canonical-2233033310311032-0112222221322210-2102221333313023-3132302233331310-1000332011200311-1201123233112013-1102213133320033-0230221020222121) |
| `items.object.spec.gc_spec.infra.hw_info.usb.port` | [items.object.spec.gc_spec.infra.hw_info.usb.port](data-sources--site_registrations_by_state--reference--group-003.md#canonical-2323202223321103-3123023101111321-0113211032301312-1322003021100023-1300131223120231-3233310332202213-3213330011303323-0113120330103000) |
| `items.object.spec.gc_spec.infra.hw_info.usb.product_name` | [items.object.spec.gc_spec.infra.hw_info.usb.product_name](data-sources--site_registrations_by_state--reference--group-003.md#canonical-1302322301003221-2010330233020210-1112230120132011-3023101031322122-3001020203133302-0022123223001012-0003232100330010-3020013022023232) |
| `items.object.spec.gc_spec.infra.hw_info.usb.speed` | [items.object.spec.gc_spec.infra.hw_info.usb.speed](data-sources--site_registrations_by_state--reference--group-003.md#canonical-2131012111222113-3203030202111210-2230123123002022-1023311311031023-1103102333231110-1101330032202020-1210312013120021-3310032030032301) |
| `items.object.spec.gc_spec.infra.hw_info.usb.usb_type` | [items.object.spec.gc_spec.infra.hw_info.usb.usb_type](data-sources--site_registrations_by_state--reference--group-003.md#canonical-3101313312233311-3001221112113332-2112303103023201-2220331221110211-0011011233232013-3311100201023101-1331212120213022-2201032313201331) |
| `items.object.spec.gc_spec.infra.hw_info.usb.vendor_name` | [items.object.spec.gc_spec.infra.hw_info.usb.vendor_name](data-sources--site_registrations_by_state--reference--group-003.md#canonical-1132231133231012-0101213303330223-2223030010322130-0220303103001102-1112130221110033-3201322110303112-3121331201022230-1132002130300301) |
| `items.object.spec.gc_spec.infra.instance_id` | [items.object.spec.gc_spec.infra.instance_id](data-sources--site_registrations_by_state--reference--group-002.md#canonical-1022103020321231-0011032023233000-3131012103022101-0022322321212300-2322200301321111-0032332022132333-0003313023131313-2213201331022133) |
| `items.object.spec.gc_spec.infra.interfaces` | [items.object.spec.gc_spec.infra.interfaces](data-sources--site_registrations_by_state--reference--group-003.md#canonical-0332303020211113-0201201320203020-3321231213030001-1001232113322000-0300021113131230-3312222302020132-0002022130231003-1310312202030033) |
| `items.object.spec.gc_spec.infra.internet_proxy` | [items.object.spec.gc_spec.infra.internet_proxy](data-sources--site_registrations_by_state--reference--group-003.md#canonical-1220332323013031-1000030211231303-1320100311130331-0130112031211202-0320302123220120-3332212211131330-0233230310313120-1130133013032210) |
| `items.object.spec.gc_spec.infra.internet_proxy.http_proxy` | [items.object.spec.gc_spec.infra.internet_proxy.http_proxy](data-sources--site_registrations_by_state--reference--group-003.md#canonical-3122213001201121-0203300131231120-0003021322230311-1012303001002310-2311113331122013-2303013020200101-3022221103030003-2002220210200322) |
| `items.object.spec.gc_spec.infra.internet_proxy.https_proxy` | [items.object.spec.gc_spec.infra.internet_proxy.https_proxy](data-sources--site_registrations_by_state--reference--group-003.md#canonical-3230102020323320-1101210013333133-1121003120131103-3110113330021021-2101233111203110-3312323331033302-3011233231030100-0310211120312222) |
| `items.object.spec.gc_spec.infra.internet_proxy.no_proxy` | [items.object.spec.gc_spec.infra.internet_proxy.no_proxy](data-sources--site_registrations_by_state--reference--group-003.md#canonical-0331021123300213-1331323020010202-3321112122000301-0330130102023022-0201221211013101-0033020220201331-3200333220033120-2112310003323230) |
| `items.object.spec.gc_spec.infra.internet_proxy.proxy_cacert_url` | [items.object.spec.gc_spec.infra.internet_proxy.proxy_cacert_url](data-sources--site_registrations_by_state--reference--group-003.md#canonical-3221223220330012-2020102200120002-3112002322132302-0200000231223120-3200103322323323-1203131333331131-0113100331011203-3221322201201001) |
| `items.object.spec.gc_spec.infra.is_slo_static` | [items.object.spec.gc_spec.infra.is_slo_static](data-sources--site_registrations_by_state--reference--group-002.md#canonical-0113003333333222-3132033331100110-2103321300313010-1133200013321111-1112233311120323-1322232012310122-1112030003221222-1313210211200231) |
| `items.object.spec.gc_spec.infra.machine_id` | [items.object.spec.gc_spec.infra.machine_id](data-sources--site_registrations_by_state--reference--group-002.md#canonical-1301033330223301-3001101003220211-0222330232002303-1320333110321130-3032310330202311-0320031013321022-3003120331020232-2330230223022110) |
| `items.object.spec.gc_spec.infra.provider_ref` | [items.object.spec.gc_spec.infra.provider_ref](data-sources--site_registrations_by_state--reference--group-002.md#canonical-3330202311233103-3113121100102032-0030330031212023-0321121232030313-3220003033222333-0021130220030231-0113322000213231-2333212221302212) |
| `items.object.spec.gc_spec.infra.sw_info` | [items.object.spec.gc_spec.infra.sw_info](data-sources--site_registrations_by_state--reference--group-003.md#canonical-3021030123331022-0012021102302222-0101010110220113-2002013122010201-2123212331233132-1232003220321022-2102323021023002-2310210311132013) |
| `items.object.spec.gc_spec.infra.sw_info.sw_version` | [items.object.spec.gc_spec.infra.sw_info.sw_version](data-sources--site_registrations_by_state--reference--group-003.md#canonical-3301230311323020-0322032220031212-2330222212111321-2211010320120330-3000311220013133-3331030223002021-3220013133003201-2311223221323301) |
| `items.object.spec.gc_spec.infra.timestamp` | [items.object.spec.gc_spec.infra.timestamp](data-sources--site_registrations_by_state--reference--group-002.md#canonical-3303111111201213-1111102212011201-1321312313012003-2233032033110300-2013220332133121-3301102313333011-3313030302112013-3020312331333222) |
| `items.object.spec.gc_spec.infra.zone` | [items.object.spec.gc_spec.infra.zone](data-sources--site_registrations_by_state--reference--group-002.md#canonical-1023332223003012-2013000202100001-2330201000011012-3102333110012101-3022232321112303-1122303210003311-3210111301011111-2030012202023231) |
| `items.object.spec.gc_spec.passport` | [items.object.spec.gc_spec.passport](data-sources--site_registrations_by_state--reference--group-003.md#canonical-3103131010303230-1121111103031012-1303120100031203-2301312000103212-1122032212110022-0030110132231123-0002111222133231-3110013232213031) |
| `items.object.spec.gc_spec.passport.cluster_name` | [items.object.spec.gc_spec.passport.cluster_name](data-sources--site_registrations_by_state--reference--group-003.md#canonical-2101103313300201-3332023120300201-1111022313102120-1110232021101231-0001302333222211-2132222211011213-1201133323110121-2302333023122122) |
| `items.object.spec.gc_spec.passport.cluster_size` | [items.object.spec.gc_spec.passport.cluster_size](data-sources--site_registrations_by_state--reference--group-003.md#canonical-2110202132302130-1212212030233130-1010201313200013-1133100013203320-1222323230100023-2322213013310221-2233131033003102-0203003000132300) |
| `items.object.spec.gc_spec.passport.cluster_type` | [items.object.spec.gc_spec.passport.cluster_type](data-sources--site_registrations_by_state--reference--group-003.md#canonical-2221011330333201-2122010001310231-3210022322132000-2130203333122311-2123111233013311-3020231133030213-0003302000020223-2102022223000331) |
| `items.object.spec.gc_spec.passport.default_os_version` | [items.object.spec.gc_spec.passport.default_os_version](data-sources--site_registrations_by_state--reference--group-003.md#canonical-3223223323202110-1220303112323030-2131023333130130-2233010113321013-0310322313302222-2003101131222233-1122023103321221-2122033132221311) |
| `items.object.spec.gc_spec.passport.default_sw_version` | [items.object.spec.gc_spec.passport.default_sw_version](data-sources--site_registrations_by_state--reference--group-003.md#canonical-2213213312031323-1031320121103321-3300020013032303-2210023202303222-1223100000212303-1131000320213111-0012010311311331-3311202300110110) |
| `items.object.spec.gc_spec.passport.latitude` | [items.object.spec.gc_spec.passport.latitude](data-sources--site_registrations_by_state--reference--group-003.md#canonical-0020222203132233-3200113121113203-3203223122222012-1231223113010202-2202322022211222-3022021111002010-2330330122023223-2211113101122312) |
| `items.object.spec.gc_spec.passport.longitude` | [items.object.spec.gc_spec.passport.longitude](data-sources--site_registrations_by_state--reference--group-003.md#canonical-0213213111033013-0311111103320303-0121203222202011-2131333132321010-0311221221332030-1303311021332033-1122211131030013-1112212213212330) |
| `items.object.spec.gc_spec.passport.operating_system_version` | [items.object.spec.gc_spec.passport.operating_system_version](data-sources--site_registrations_by_state--reference--group-003.md#canonical-1033123331333230-2203020200100031-1301000300231110-2200213103310033-1200120203032022-1101300210231322-2223131102133311-0301320110010133) |
| `items.object.spec.gc_spec.passport.private_network_name` | [items.object.spec.gc_spec.passport.private_network_name](data-sources--site_registrations_by_state--reference--group-003.md#canonical-0002203020000230-1023213221222133-0023113332202001-0321000301223130-2203302330000323-1301233033021320-3230311200311023-3301031310003332) |
| `items.object.spec.gc_spec.passport.volterra_software_version` | [items.object.spec.gc_spec.passport.volterra_software_version](data-sources--site_registrations_by_state--reference--group-003.md#canonical-0021003111100210-2330032013210003-0112112032300000-0113322000123200-2021100031023201-2312031002301010-2201230222010223-1203300020021013) |
| `items.object.spec.gc_spec.role` | [items.object.spec.gc_spec.role](data-sources--site_registrations_by_state--reference--group-002.md#canonical-1021022121021322-0323031303333133-3020012110212013-0201333020123311-3203132130233332-0310130113202321-0301202301231323-2030103301213133) |
| `items.object.spec.gc_spec.site` | [items.object.spec.gc_spec.site](data-sources--site_registrations_by_state--reference--group-003.md#canonical-2111301010302203-2121000033123223-2033303313102311-2203030321001220-2333201201301202-0223331132323102-1220032121210310-3331333300323202) |
| `items.object.spec.gc_spec.site.kind` | [items.object.spec.gc_spec.site.kind](data-sources--site_registrations_by_state--reference--group-003.md#canonical-2201021300210313-2201112310132223-2221212233030233-0123030112132310-3201032021220331-2211110323000220-1210000112321233-2122232102103110) |
| `items.object.spec.gc_spec.site.name` | [items.object.spec.gc_spec.site.name](data-sources--site_registrations_by_state--reference--group-003.md#canonical-2103013303031022-2131122022201001-3100303211030023-2212121022010211-1320230002310112-2010203032132231-1100312001013300-3011103223212010) |
| `items.object.spec.gc_spec.site.namespace` | [items.object.spec.gc_spec.site.namespace](data-sources--site_registrations_by_state--reference--group-003.md#canonical-2212203222001210-0021101200323223-2010120123330031-1313111223310302-2230303030003202-0100223320121012-1211333132020311-1222010123133103) |
| `items.object.spec.gc_spec.site.tenant` | [items.object.spec.gc_spec.site.tenant](data-sources--site_registrations_by_state--reference--group-003.md#canonical-1213122221032002-0030031321313322-2113203330020011-2203303030032321-0032130012010331-1313022022023300-1103332102313212-1011022300120023) |
| `items.object.spec.gc_spec.site.uid` | [items.object.spec.gc_spec.site.uid](data-sources--site_registrations_by_state--reference--group-003.md#canonical-0122131101133200-1123010201231023-1300211123213230-0102010030000313-3222020112021002-2232203023030313-2311001113110331-0112023012100100) |
| `items.object.spec.gc_spec.token` | [items.object.spec.gc_spec.token](data-sources--site_registrations_by_state--reference--group-002.md#canonical-1121231311110221-3312112333032301-0103001011320223-0000202222333302-3132300021032001-3130011320020233-2030322233121213-1113100112223320) |
| `items.object.spec.gc_spec.tunnel_type` | [items.object.spec.gc_spec.tunnel_type](data-sources--site_registrations_by_state--reference--group-002.md#canonical-1212111201311230-3312200303200123-1110301021203131-1013303032110012-3203310023023311-1223322301321112-2132301101102303-2121021333121330) |
| `items.object.status` | [items.object.status](data-sources--site_registrations_by_state--reference--group-003.md#canonical-1200223300332122-0012203011200332-0110220132030230-2000121032010121-2203210213302302-3312201311113311-2012300313012312-1120030213302213) |
| `items.object.status.current_state` | [items.object.status.current_state](data-sources--site_registrations_by_state--reference--group-003.md#canonical-3102130113030102-3303300332001130-3022220303030311-2001302200121200-2013321323300130-1320231123233131-0300131233101001-2101030033022111) |
| `items.object.status.object_status` | [items.object.status.object_status](data-sources--site_registrations_by_state--reference--group-003.md#canonical-1001100122331201-2201212230202121-2210233110013103-2331112020203223-1233012213311033-0220102320230113-1201103122330220-2332100022010103) |
| `items.object.status.object_status.code` | [items.object.status.object_status.code](data-sources--site_registrations_by_state--reference--group-003.md#canonical-2322021010032321-2311223231113331-3303011230332101-1212210113012233-2022222033223000-0130030130112020-1211322301011020-3203033031032300) |
| `items.object.status.object_status.reason` | [items.object.status.object_status.reason](data-sources--site_registrations_by_state--reference--group-003.md#canonical-2212332230001133-1200332023120133-2020323010123231-1123131211300321-1221102201101021-0313313103112313-0313221000131233-0022202233311111) |
| `items.object.status.object_status.status` | [items.object.status.object_status.status](data-sources--site_registrations_by_state--reference--group-003.md#canonical-0132033201233113-1123123222323210-0202121122230003-2321033103100012-3331032132131232-0002332231213033-0131301302313103-1103333210112302) |
| `items.object.status.parent_current_state` | [items.object.status.parent_current_state](data-sources--site_registrations_by_state--reference--group-003.md#canonical-0030222331323101-0312032333310011-0020021103230030-1220021103330223-0212310033212120-0012032003222031-0011132220231211-3003202000231013) |
| `items.object.status.state_update_timestamp` | [items.object.status.state_update_timestamp](data-sources--site_registrations_by_state--reference--group-003.md#canonical-1033221133113232-1312032221032222-3220303200320231-3102332332001001-0233100230312333-1102001103201223-2313131210102323-2303022332333211) |
| `items.object.system_metadata` | [items.object.system_metadata](data-sources--site_registrations_by_state--reference--group-003.md#canonical-0333313131030121-2003220011021132-1333003230123203-1320113310001323-1332023301301101-2301203230010311-0332000311320033-3213232030110301) |
| `items.object.system_metadata.creation_timestamp` | [items.object.system_metadata.creation_timestamp](data-sources--site_registrations_by_state--reference--group-003.md#canonical-3301012322033132-3213220122313222-2212313221201232-2313323002102302-0313010133213032-1102113232130333-3010020123302110-0331332313332132) |
| `items.object.system_metadata.creator_class` | [items.object.system_metadata.creator_class](data-sources--site_registrations_by_state--reference--group-003.md#canonical-2303312022032000-2211131120020200-1121232330101203-0230123300213132-0103331030202003-0313123330323113-1232333023023302-2033220021212001) |
| `items.object.system_metadata.creator_cookie` | [items.object.system_metadata.creator_cookie](data-sources--site_registrations_by_state--reference--group-003.md#canonical-3130032232201301-2030023131333312-0201312230022300-0132332113100133-1011113300212023-2201330312220222-2032012021131012-3333020002213110) |
| `items.object.system_metadata.creator_id` | [items.object.system_metadata.creator_id](data-sources--site_registrations_by_state--reference--group-003.md#canonical-2121002333300131-2032332010130030-0302122001330203-3213231120033123-0321033013111121-0122323222232022-3110103320113112-1110300110021011) |
| `items.object.system_metadata.deletion_timestamp` | [items.object.system_metadata.deletion_timestamp](data-sources--site_registrations_by_state--reference--group-003.md#canonical-3200331132230302-1331111103222103-3332301311212021-0030320223220102-3310133323200021-3013232012032120-3113320231000120-1002100210121213) |
| `items.object.system_metadata.direct_ref_hash` | [items.object.system_metadata.direct_ref_hash](data-sources--site_registrations_by_state--reference--group-003.md#canonical-2001002333121132-0320131120310122-2200213033303113-1001130211112133-2212121322223010-2211332110131300-3300010221132001-2020230202322032) |
| `items.object.system_metadata.finalizers` | [items.object.system_metadata.finalizers](data-sources--site_registrations_by_state--reference--group-003.md#canonical-2212322320222330-2230111333100310-3202013203113212-2030013220221231-0132100323012223-2002132032012210-1212233030202322-3302013211331232) |
| `items.object.system_metadata.initializers` | [items.object.system_metadata.initializers](data-sources--site_registrations_by_state--reference--group-003.md#canonical-3101020020131332-3223332010320203-1322233032311313-0020321112102031-0200121301331033-1130300223013210-3120220312112000-2102230330001312) |
| `items.object.system_metadata.initializers.pending` | [items.object.system_metadata.initializers.pending](data-sources--site_registrations_by_state--reference--group-003.md#canonical-3213021221221330-2103312233120311-0031313310010020-1000200202130011-2003030222131032-0001132211303201-1210232312113233-0232232220120011) |
| `items.object.system_metadata.initializers.pending.name` | [items.object.system_metadata.initializers.pending.name](data-sources--site_registrations_by_state--reference--group-003.md#canonical-0213333023113313-1321221003133331-1231200303212033-2321222112031032-0112322113300213-0011103200133031-1312223120102113-2002203020302332) |
| `items.object.system_metadata.initializers.result` | [items.object.system_metadata.initializers.result](data-sources--site_registrations_by_state--reference--group-003.md#canonical-1000230231020222-2211111011233013-3231230112133111-2033333121330113-1303101111331120-0301203203332132-2002112222021210-2121020321331031) |
| `items.object.system_metadata.initializers.result.code` | [items.object.system_metadata.initializers.result.code](data-sources--site_registrations_by_state--reference--group-003.md#canonical-3311120133223110-0112232303331332-0023233303231310-0111320023200201-1301022301220212-1111031022112000-1221211010233300-1301211030023001) |
| `items.object.system_metadata.initializers.result.reason` | [items.object.system_metadata.initializers.result.reason](data-sources--site_registrations_by_state--reference--group-003.md#canonical-1333320023032122-2320102311331001-0312022210302100-0103310122331001-0103323101113111-3010123222213010-2030010112211012-2233322231130320) |
| `items.object.system_metadata.initializers.result.status` | [items.object.system_metadata.initializers.result.status](data-sources--site_registrations_by_state--reference--group-003.md#canonical-1231010032001200-3202012223010323-3111003102211112-3303222202332320-3111230020103203-0000021130300320-0221222003100031-0001313000210313) |
| `items.object.system_metadata.labels` | [items.object.system_metadata.labels](data-sources--site_registrations_by_state--reference--group-003.md#canonical-0023302321310211-3011123321122313-3132122202001322-0022112001133012-0100001231322332-1230132231301032-1101301320003211-1030033311133002) |
| `items.object.system_metadata.modification_timestamp` | [items.object.system_metadata.modification_timestamp](data-sources--site_registrations_by_state--reference--group-003.md#canonical-2212133031211303-2200202121010002-2330200123230211-2203130232313103-3213330221330313-3001211210310232-0312333321103013-3002111203133201) |
| `items.object.system_metadata.namespace` | [items.object.system_metadata.namespace](data-sources--site_registrations_by_state--reference--group-003.md#canonical-0012222222021000-0000013121232003-1321130111111302-3222211221231321-3120323310011223-3202100013133201-2313202120210023-1212102020213012) |
| `items.object.system_metadata.namespace.kind` | [items.object.system_metadata.namespace.kind](data-sources--site_registrations_by_state--reference--group-003.md#canonical-1032121021033030-0133212121332130-0200221310031202-2100320121111310-0201130222002222-3010003130011203-1310032333100220-0323111130002121) |
| `items.object.system_metadata.namespace.name` | [items.object.system_metadata.namespace.name](data-sources--site_registrations_by_state--reference--group-003.md#canonical-3021002101201112-2132011022201122-0222333130321221-3121212002223203-2011213333113112-3210221113213013-2230101313112031-1320030023031211) |
| `items.object.system_metadata.namespace.namespace` | [items.object.system_metadata.namespace.namespace](data-sources--site_registrations_by_state--reference--group-003.md#canonical-3301302122102200-1331210012211223-0321003322011123-0003103222110231-0221302103112010-1101012023021133-0100133223230131-0320301002001100) |
| `items.object.system_metadata.namespace.tenant` | [items.object.system_metadata.namespace.tenant](data-sources--site_registrations_by_state--reference--group-003.md#canonical-0320220202021233-0113123301122222-2301030030321221-3000222012201230-2331313320232300-0110010120212031-2301211310031031-0232001023212223) |
| `items.object.system_metadata.namespace.uid` | [items.object.system_metadata.namespace.uid](data-sources--site_registrations_by_state--reference--group-003.md#canonical-0121031020232330-0223323023231210-1002330303031302-2322132302223121-1130103220011120-3222322302302001-3203112312031033-3200010113133222) |
| `items.object.system_metadata.object_index` | [items.object.system_metadata.object_index](data-sources--site_registrations_by_state--reference--group-003.md#canonical-3233321012321311-2113033212131021-0110120231120223-2300332221300112-0000002312202212-1333120013301100-0002323222022213-1231323011313203) |
| `items.object.system_metadata.owner_view` | [items.object.system_metadata.owner_view](data-sources--site_registrations_by_state--reference--group-003.md#canonical-2200321323100031-0313313121200023-0032133001333102-0123021110221231-1022011100101033-3101332001213203-2112021210201201-2300303022222101) |
| `items.object.system_metadata.owner_view.kind` | [items.object.system_metadata.owner_view.kind](data-sources--site_registrations_by_state--reference--group-003.md#canonical-2011202133022030-2020031211112032-0223022330030300-0222320123313101-0200212110210222-1022331123220102-0122033033100120-1300023010113003) |
| `items.object.system_metadata.owner_view.name` | [items.object.system_metadata.owner_view.name](data-sources--site_registrations_by_state--reference--group-003.md#canonical-2023100000223020-1321231111002123-2100001121230001-2310103022012311-1133302323200321-1002313221100201-2313031303102333-2333002333113223) |
| `items.object.system_metadata.owner_view.namespace` | [items.object.system_metadata.owner_view.namespace](data-sources--site_registrations_by_state--reference--group-003.md#canonical-0101020032221113-0010222031113300-2021101031303220-2032222121213302-0001203113000011-2202312022011003-1303232230130311-2201111033232213) |
| `items.object.system_metadata.owner_view.uid` | [items.object.system_metadata.owner_view.uid](data-sources--site_registrations_by_state--reference--group-003.md#canonical-1010312311312123-0033201303332310-1120310300230111-0221311321122133-2132110102231233-1321032303030331-1000231323110030-3200303110231302) |
| `items.object.system_metadata.revision` | [items.object.system_metadata.revision](data-sources--site_registrations_by_state--reference--group-003.md#canonical-3201103200201210-3310100230332203-1021212003012031-1220120031331002-0033101323333131-0123123211111223-3112103102130332-0230003011101231) |
| `items.object.system_metadata.sre_disable` | [items.object.system_metadata.sre_disable](data-sources--site_registrations_by_state--reference--group-003.md#canonical-3220230301231211-1232032222123213-3312302200013202-1001232312112022-0010301232032021-3112202331300020-0222222123331020-0210313031100111) |
| `items.object.system_metadata.tenant` | [items.object.system_metadata.tenant](data-sources--site_registrations_by_state--reference--group-003.md#canonical-0301320323211200-1022312210332311-3331013123021230-0232321122032202-1101200102020032-0122032111003222-1030222301110202-3111203103001100) |
| `items.object.system_metadata.trace_info` | [items.object.system_metadata.trace_info](data-sources--site_registrations_by_state--reference--group-003.md#canonical-0322111213200010-2113233231333033-2122122101031202-2000332000102033-3121223213111023-0030111021220103-2021312033011231-3203022100112101) |
| `items.object.system_metadata.uid` | [items.object.system_metadata.uid](data-sources--site_registrations_by_state--reference--group-003.md#canonical-2213231332212032-0311111102212031-1131021232201313-1233221133112033-3320032300323231-2012231331031223-0211322110012203-3013302010001321) |
| `items.object.system_metadata.vtrp_id` | [items.object.system_metadata.vtrp_id](data-sources--site_registrations_by_state--reference--group-003.md#canonical-0211303002111222-0330230000130213-0220301031322210-3100333310102322-2101303112121011-0210001300031102-3002202121222330-3300330331331123) |
| `items.object.system_metadata.vtrp_stale` | [items.object.system_metadata.vtrp_stale](data-sources--site_registrations_by_state--reference--group-003.md#canonical-1333313213222321-2120231031312013-0323310330213010-0110221302213022-3320202320313122-1230311121030203-0132200200233203-2331101130303223) |
| `items.owner_view` | [items.owner_view](data-sources--site_registrations_by_state--reference--group-003.md#canonical-3130223130131223-3133330130032303-1101002223320112-2302213112203031-1200123212001213-1100011131103212-1002212230113113-0113002233311232) |
| `items.owner_view.kind` | [items.owner_view.kind](data-sources--site_registrations_by_state--reference--group-003.md#canonical-1210002321211111-2000113131023202-2133021311212331-1031102221101302-3310102031313320-3122032032310200-3020031333332031-1202231003100320) |
| `items.owner_view.name` | [items.owner_view.name](data-sources--site_registrations_by_state--reference--group-003.md#canonical-0001112301331003-1201230010311111-3331112002312221-0311300300211123-0101013113113131-1132120023031122-3312312231001303-2210310230313320) |
| `items.owner_view.namespace` | [items.owner_view.namespace](data-sources--site_registrations_by_state--reference--group-003.md#canonical-1231233100202311-3001110102302031-1323203223222333-3112330213132122-2303303100111031-1322133100111020-0203003001133310-0030113130213112) |
| `items.owner_view.uid` | [items.owner_view.uid](data-sources--site_registrations_by_state--reference--group-003.md#canonical-2333122313333220-2202003323132311-0032012120212000-3120101311321331-3211221313022130-3222122302131311-2103021203033021-3021000132233200) |
| `items.system_metadata` | [items.system_metadata](data-sources--site_registrations_by_state--reference--group-003.md#canonical-3323332131030132-0102131330332310-1313133111033131-3033012332202011-3102021301311330-0131122013033033-3200122321223032-0311032013122020) |
| `items.system_metadata.creation_timestamp` | [items.system_metadata.creation_timestamp](data-sources--site_registrations_by_state--reference--group-003.md#canonical-3101111222310012-1222000303110201-1123030323332000-3101000321300233-3103232210011231-3131120232202200-1023031332130302-1230222112320121) |
| `items.system_metadata.creator_class` | [items.system_metadata.creator_class](data-sources--site_registrations_by_state--reference--group-003.md#canonical-3123321332121133-3023301123200231-2210121231221302-3202210103023310-0310322003002033-2000033032212011-3223013001330102-1311012200131130) |
| `items.system_metadata.creator_id` | [items.system_metadata.creator_id](data-sources--site_registrations_by_state--reference--group-003.md#canonical-0301323230230120-3103130132313301-0003332220233111-1201333222333212-3223032112203031-0213202330031233-2001100113030101-0133102333121022) |
| `items.system_metadata.deletion_timestamp` | [items.system_metadata.deletion_timestamp](data-sources--site_registrations_by_state--reference--group-003.md#canonical-3321122021211303-3113222301013222-0202322103232110-3002333112213331-1030313030223022-2100210132200302-2022102302002330-1021231001033303) |
| `items.system_metadata.finalizers` | [items.system_metadata.finalizers](data-sources--site_registrations_by_state--reference--group-003.md#canonical-3223130303300200-3303321020332303-3113100130120321-0233131031312030-1010133320100121-0021333132230313-2213033201100313-2331210213203230) |
| `items.system_metadata.initializers` | [items.system_metadata.initializers](data-sources--site_registrations_by_state--reference--group-003.md#canonical-3200333333111332-3003331011223203-1133023333202021-1102202123333330-1110022232103221-2331032001210231-0010333132201102-2003030320133232) |
| `items.system_metadata.initializers.pending` | [items.system_metadata.initializers.pending](data-sources--site_registrations_by_state--reference--group-003.md#canonical-0212021200120302-0021211213333203-2331220302231211-0211031111302230-1130231332201003-1220130131132102-3030110322133231-2232032301213213) |
| `items.system_metadata.initializers.pending.name` | [items.system_metadata.initializers.pending.name](data-sources--site_registrations_by_state--reference--group-003.md#canonical-1111232033222301-1100231320102001-1330331031011212-0102212210100110-1323213211212100-2013213001120031-0132101001003230-1210203032203010) |
| `items.system_metadata.initializers.result` | [items.system_metadata.initializers.result](data-sources--site_registrations_by_state--reference--group-003.md#canonical-3213121031221102-3211002010203233-2323012223110220-0200233233303120-2032233030230021-1001123220210220-3313122011032111-1312210213031132) |
| `items.system_metadata.initializers.result.code` | [items.system_metadata.initializers.result.code](data-sources--site_registrations_by_state--reference--group-003.md#canonical-1033013132300212-1220011001310200-3133010221311031-2103120331301232-3212201132201330-3003031203111113-2323030202232212-2321200200031210) |
| `items.system_metadata.initializers.result.reason` | [items.system_metadata.initializers.result.reason](data-sources--site_registrations_by_state--reference--group-003.md#canonical-1100001130223302-0033222131320011-3032121033300231-2233332031130330-1312032300212123-2033002213113121-2022330023101033-3111113223003113) |
| `items.system_metadata.initializers.result.status` | [items.system_metadata.initializers.result.status](data-sources--site_registrations_by_state--reference--group-003.md#canonical-3210103312023332-2231312310100311-1112211210022233-2333130113303300-1032120013203200-0312000220002132-1101321100131103-0233111212332021) |
| `items.system_metadata.labels` | [items.system_metadata.labels](data-sources--site_registrations_by_state--reference--group-003.md#canonical-2120012202021200-3033032123030323-3310110221322201-3303311011310123-0001023002133301-2111201313121000-3101022332333033-1120123300300232) |
| `items.system_metadata.modification_timestamp` | [items.system_metadata.modification_timestamp](data-sources--site_registrations_by_state--reference--group-003.md#canonical-0012202332120020-0212222013032313-3323311130111213-1300300203023202-3312331031303003-0130101011213323-2313113130110002-1201101031122000) |
| `items.system_metadata.object_index` | [items.system_metadata.object_index](data-sources--site_registrations_by_state--reference--group-003.md#canonical-2122012232000231-3013212030212211-2132003233031230-0132020123021023-3133002120313031-0220111230132320-0221200223113331-3202203133122030) |
| `items.system_metadata.owner_view` | [items.system_metadata.owner_view](data-sources--site_registrations_by_state--reference--group-003.md#canonical-0232131232210230-3202301213110003-3012323211202312-3010302332222332-0222312000102231-1023033320303022-1333031321112033-0002013112310110) |
| `items.system_metadata.owner_view.kind` | [items.system_metadata.owner_view.kind](data-sources--site_registrations_by_state--reference--group-003.md#canonical-3123231301003002-3201201223323231-0031031133320132-2320303302012231-1010323220201013-0312321022122210-0113010103212102-0200120330131113) |
| `items.system_metadata.owner_view.name` | [items.system_metadata.owner_view.name](data-sources--site_registrations_by_state--reference--group-003.md#canonical-1132221130010111-3000313032300010-3003010201032323-3020022322132131-1231313000001232-0120020102132021-1332131313130033-0230002112021230) |
| `items.system_metadata.owner_view.namespace` | [items.system_metadata.owner_view.namespace](data-sources--site_registrations_by_state--reference--group-003.md#canonical-3303022101222213-3202202101000100-3112203201231133-2212103111302112-3132210221133012-0203332100330110-3120111102223203-1312002221133111) |
| `items.system_metadata.owner_view.uid` | [items.system_metadata.owner_view.uid](data-sources--site_registrations_by_state--reference--group-003.md#canonical-0210001312123203-3303301031010111-1102030212021002-2311021230013101-0121111232102003-3230310203323122-1030300001130312-3021233013313121) |
| `items.system_metadata.tenant` | [items.system_metadata.tenant](data-sources--site_registrations_by_state--reference--group-003.md#canonical-1112031300130311-0221113110020222-2132300320010321-1230310213132202-1312311033000223-3032101313012201-2010120220110030-2122130203120233) |
| `items.system_metadata.uid` | [items.system_metadata.uid](data-sources--site_registrations_by_state--reference--group-003.md#canonical-0332121201030202-3311122020333313-0100301101120100-1331202310220320-0210103013323013-3132003231203011-3220213132200013-0233022230213233) |
| `items.tenant` | [items.tenant](data-sources--site_registrations_by_state--reference--group-001.md#canonical-3330132121301300-2312302100130030-3133131320300102-0100210120010110-2103210112323223-3211313201033123-0232320312023031-0122101233120201) |
| `items.uid` | [items.uid](data-sources--site_registrations_by_state--reference--group-001.md#canonical-2332122233111233-2210020323021233-3201131300013232-1122223221120302-2123200001131213-0012301133213302-2233123002131332-3202010010121103) |
| `namespace` | [namespace](data-sources--site_registrations_by_state--reference--group-001.md#canonical-1311001021113033-0013020103210231-0132213023213212-0331132313231220-3211102013021322-0023032320202320-2231220311230021-3332220303110233) |
| `state` | [state](data-sources--site_registrations_by_state--reference--group-001.md#canonical-3131223012333110-0130012003031211-1120300301022230-1000230123130111-2320100001031112-0231313202200313-3123133123000213-0000330130133310) |

<a id="canonical-3123201032113231-1110203102030102-0321102211032320-2011122333202121-2011102030213113-2300113212013233-3311310101312121-2030132221300302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `errors` properties

Breadcrumbs:

- [xcsh_site_registrations_by_state](../data-sources/site_registrations_by_state.md#canonical-1011310232230021-3303123312003111-3120023223111201-0320322111010320-3032200132321200-3310101021222131-0120032131300331-0201133223313303)
- [Property reference](data-sources--site_registrations_by_state--reference--group-001.md#canonical-0320011213231301-1102033233211220-0302131120233130-2122121213220203-2033211132132211-0333231313221113-3210300123023021-2021122221230130)
- errors

<a id="canonical-0233103032201323-0133210100320323-0123313010223300-2222021111020023-0103110100010333-0323130331130120-3020300121132130-2332200221113102"></a>

Type: `"list"`. Computed.

Errors(if any) while listing items from collection.

<a id="canonical-2011231201013333-2331231313021301-3232230010220312-0011200112000232-1113332210122003-0121132010222130-0002332221322002-3213033033211312"></a>

### Direct properties for `errors`

<a id="canonical-0211303333302112-2032313300321130-1233002311031101-2132030133232303-0212310130021323-0322222132210313-2211310033130121-1002132030020322"></a>

#### `errors.code` property

Type: `"string"`. Computed.

\[Enum: EOK|EPERMS|EBADINPUT|ENOTFOUND|EEXISTS|EUNKNOWN|ESERIALIZE|EINTERNAL|EPARTIAL\] Union of all
possible error-codes from system - EOK: No error - EPERMS: Permissions error - EBADINPUT: Input is
not correct - ENOTFOUND: Not found - EEXISTS: Already exists - EUNKNOWN: Unknown/catchall error -
ESERIALIZE: Error in serializing/de-serializing - EINTERNAL: Server error - EPARTIAL.. Possible
values are \`EOK\`, \`EPERMS\`, \`EBADINPUT\`, \`ENOTFOUND\`, \`EEXISTS\`, \`EUNKNOWN\`,
\`ESERIALIZE\`, \`EINTERNAL\`, \`EPARTIAL\`. Defaults to \`EOK\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["EBADINPUT","EEXISTS","EINTERNAL","ENOTFOUND","EOK","EPARTIAL","EPERMS","ESERIALIZE","EUNKNOWN"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("EOK",
    "EPERMS",
    "EBADINPUT",
    "ENOTFOUND",
    "EEXISTS",
    "EUNKNOWN",
    "ESERIALIZE",
    "EINTERNAL",
    "EPARTIAL"),
}
```

- [error_obj](data-sources--site_registrations_by_state--reference--group-001.md#canonical-1010211013112301-1131221010330231-3133220301030231-0302212300122302-3200131000020122-0110213101233100-3120011230322302-0313320102333033): complete subsection reference.

<a id="canonical-0102232032001010-1330123103321012-0020230300023010-2311330130311301-3221201331021212-0101323331313213-1021311321330002-2212030300022210"></a>

<a id="canonical-0212120232031000-1132120331330000-1012000112012332-0013131310031133-2023213323031033-1001033333132233-1020221130322121-1330313031112002"></a>

#### `errors.message` property

Type: `"string"`. Computed.

Message. A human readable string of the error.

<a id="canonical-1010211013112301-1131221010330231-3133220301030231-0302212300122302-3200131000020122-0110213101233100-3120011230322302-0313320102333033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `errors.error_obj` properties

Breadcrumbs:

- [xcsh_site_registrations_by_state](../data-sources/site_registrations_by_state.md#canonical-1011310232230021-3303123312003111-3120023223111201-0320322111010320-3032200132321200-3310101021222131-0120032131300331-0201133223313303)
- [Property reference](data-sources--site_registrations_by_state--reference--group-001.md#canonical-0320011213231301-1102033233211220-0302131120233130-2122121213220203-2033211132132211-0333231313221113-3210300123023021-2021122221230130)
- [errors](data-sources--site_registrations_by_state--reference--group-001.md#canonical-3123201032113231-1110203102030102-0321102211032320-2011122333202121-2011102030213113-2300113212013233-3311310101312121-2030132221300302)
- errors.error_obj

<a id="canonical-2022310132031021-3031031023021221-0113302330030003-3100223213020000-2201312000111122-0021320020103322-3032311302020102-3321301300120220"></a>

Type: `"single"`. Computed.

Contains an arbitrary serialized protocol buffer message along with a URL that describes the type of
the serialized message. Protobuf library provides support to pack/unpack Any values in the form of
utility functions or additional generated methods of the Any type. Example 1: Pack and unpack a..

<a id="canonical-2322321212131230-0102230323323132-2133020322100111-0031032002120230-2310233132210110-3320032203202013-1113220233023321-1111022132333113"></a>

### Direct properties for `errors.error_obj`

<a id="canonical-3010100033003310-1133011021213222-2331210301103300-1333232130123233-1200020132302333-2221132133223330-0002221120130303-1032310310333031"></a>

#### `errors.error_obj.type_url` property

Type: `"string"`. Computed.

URL/resource name that uniquely identifies the type of the serialized protocol buffer message. This
string must contain at least one '/' character. The last segment of the URL path must represent the
fully qualified name of the type (as in ).

<a id="canonical-1000022001002300-0003223321311221-3320312021101131-2031022232310101-1030210320100023-0033222210203033-3231311220213332-1220022010331133"></a>

<a id="canonical-2110223232321210-1132112233010322-1101010331303301-1323031323302002-0211120233332120-0301031200101000-0112201022033312-0021313132100122"></a>

#### `errors.error_obj.value` property

Type: `"string"`. Computed.

Must be a valid serialized protocol buffer of the above specified type.

<a id="canonical-1133122031213013-2312233003121321-2003030111222220-0123212333122102-1113033120112122-1330232222330233-1111303223103223-2020000130121230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `items` properties

Breadcrumbs:

- [xcsh_site_registrations_by_state](../data-sources/site_registrations_by_state.md#canonical-1011310232230021-3303123312003111-3120023223111201-0320322111010320-3032200132321200-3310101021222131-0120032131300331-0201133223313303)
- [Property reference](data-sources--site_registrations_by_state--reference--group-001.md#canonical-0320011213231301-1102033233211220-0302131120233130-2122121213220203-2033211132132211-0333231313221113-3210300123023021-2021122221230130)
- items

<a id="canonical-2323020000030220-2023122022323311-1120031030030330-1311031111231111-3102210202310313-1111311200122133-3103213332021312-0033132021100303"></a>

Type: `"list"`. Computed.

Items represents the collection in response.

<a id="canonical-1001221122312312-3330231000010221-3303103113311121-0202103301201121-3130003230102232-3201103302001103-1321121100000013-2130010001123220"></a>

### Direct properties for `items`

- [annotations](data-sources--site_registrations_by_state--reference--group-001.md#canonical-1111112331011032-2101233230200031-2220220221020023-3321111201233102-0223101023313321-1233013301230120-1002103232202012-1101001101021030): complete subsection reference.

<a id="canonical-3233022003233203-3022033321311213-1323020213130021-0021212211033020-2331010022330303-3202232232031003-3320022221110200-3122022230100030"></a>

<a id="canonical-3022102023232323-0120102331202322-2111021102023021-0210133310013012-0101020313313031-1312323020232310-3332030000330012-2222011020100320"></a>

#### `items.description_spec` property

Type: `"string"`. Computed.

The description set for this registration.

<a id="canonical-0103322331132121-3233010012332023-1131111122300001-0211111131332133-2011112332130210-3100221002202322-2022210323011132-3100322033010220"></a>

<a id="canonical-2132210232312203-2000301111010301-0003020201020121-1202223202030112-1032132121123121-2212201110110012-0121220332300130-3113011231132000"></a>

#### `items.disabled` property

Type: `"bool"`. Computed.

Value of true indicates registration is administratively disabled.

- [get_spec](data-sources--site_registrations_by_state--reference--group-001.md#canonical-2321111233223323-1103131313132111-1332011333123121-2230221231113011-3312330012020320-0211132223121102-2102332200012230-0013300300311322): complete subsection reference.

- [labels](data-sources--site_registrations_by_state--reference--group-002.md#canonical-0113333120313032-0333021001310200-2302201021223222-1013131210011232-1333103223220323-3110321120132032-2312201232023001-2001003020222102): complete subsection reference.

- [metadata](data-sources--site_registrations_by_state--reference--group-002.md#canonical-0002032302103311-0221203213023231-3122021003312130-1331101011003121-2323222002112200-3002012331110331-1321330031023123-2013130031230103): complete subsection reference.

<a id="canonical-1021003023132032-1011113231020112-0031333132102310-1032320201201131-1213010221023301-0312113013032220-0210313210000120-2203320330312000"></a>

<a id="canonical-1031321113101231-3002011102223201-3331323112103033-3233303113312101-1322130300011013-0201120012310322-1200032001200101-0233132032230110"></a>

#### `items.name` property

Type: `"string"`. Computed.

Name. The name of this registration.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

<a id="canonical-3221000022100322-1203121313020203-1013301011022202-1123332321331130-0230322023303311-0202201202312200-1202231300203333-0213003133130233"></a>

<a id="canonical-3000120123022002-3201203320323331-0323200112032022-1020002310133123-0121313302212010-0321001211211022-1030020323011030-2333230213203333"></a>

#### `items.namespace` property

Type: `"string"`. Computed.

Namespace. The namespace this item belongs to.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

- [object](data-sources--site_registrations_by_state--reference--group-002.md#canonical-1223022213213010-2123312320000202-3212010203300330-2123121110030020-2121322322103003-3031303200312203-3323112222021022-0003011010032003): complete subsection reference.

- [owner_view](data-sources--site_registrations_by_state--reference--group-003.md#canonical-2133013002301323-3123031212331031-1032210133102130-1010222011030233-0130032223321211-2221211010323312-3103211001231130-0122303113222322): complete subsection reference.

- [system_metadata](data-sources--site_registrations_by_state--reference--group-003.md#canonical-0322321200130311-2201333230012031-1313300210121321-2003310001233032-0002023213213302-3011322012112013-1023203013031122-0110211213032322): complete subsection reference.

<a id="canonical-3330132121301300-2312302100130030-3133131320300102-0100210120010110-2103210112323223-3211313201033123-0232320312023031-0122101233120201"></a>

<a id="canonical-2002103213310111-0212010303111311-1103223111210033-1012332221103332-3130233233121113-2233300111130111-1212132303213032-0311313003002101"></a>

#### `items.tenant` property

Type: `"string"`. Computed.

Tenant. The tenant this item belongs to.

<a id="canonical-2332122233111233-2210020323021233-3201131300013232-1122223221120302-2123200001131213-0012301133213302-2233123002131332-3202010010121103"></a>

<a id="canonical-3322113120230232-2120122031013033-1023010121003101-0323101003021331-1333122022213121-0102031020122120-3000130300300321-2221033103031022"></a>

#### `items.uid` property

Type: `"string"`. Computed.

UID. The unique uid of this registration.

<a id="canonical-1111112331011032-2101233230200031-2220220221020023-3321111201233102-0223101023313321-1233013301230120-1002103232202012-1101001101021030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `items.annotations` properties

Breadcrumbs:

- [xcsh_site_registrations_by_state](../data-sources/site_registrations_by_state.md#canonical-1011310232230021-3303123312003111-3120023223111201-0320322111010320-3032200132321200-3310101021222131-0120032131300331-0201133223313303)
- [Property reference](data-sources--site_registrations_by_state--reference--group-001.md#canonical-0320011213231301-1102033233211220-0302131120233130-2122121213220203-2033211132132211-0333231313221113-3210300123023021-2021122221230130)
- [items](data-sources--site_registrations_by_state--reference--group-001.md#canonical-1133122031213013-2312233003121321-2003030111222220-0123212333122102-1113033120112122-1330232222330233-1111303223103223-2020000130121230)
- items.annotations

<a id="canonical-0332101012313233-0300112001301102-3313112333203302-3121131130000223-2232202213211012-2132101110212010-0122232103331120-2220132020011100"></a>

Type: `"single"`. Computed.

The set of annotations present on this registration.

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2321111233223323-1103131313132111-1332011333123121-2230221231113011-3312330012020320-0211132223121102-2102332200012230-0013300300311322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `items.get_spec` properties

Breadcrumbs:

- [xcsh_site_registrations_by_state](../data-sources/site_registrations_by_state.md#canonical-1011310232230021-3303123312003111-3120023223111201-0320322111010320-3032200132321200-3310101021222131-0120032131300331-0201133223313303)
- [Property reference](data-sources--site_registrations_by_state--reference--group-001.md#canonical-0320011213231301-1102033233211220-0302131120233130-2122121213220203-2033211132132211-0333231313221113-3210300123023021-2021122221230130)
- [items](data-sources--site_registrations_by_state--reference--group-001.md#canonical-1133122031213013-2312233003121321-2003030111222220-0123212333122102-1113033120112122-1330232222330233-1111303223103223-2020000130121230)
- items.get_spec

<a id="canonical-2212013122131122-3133132230212122-2001122231031230-2103210323131320-3230030333201001-3001231030012030-2121321103212133-0002002113203330"></a>

Type: `"single"`. Computed.

GET Registration. GET registration specification.

<a id="canonical-0322010210111001-3022222110312132-3230121111031133-3002311033123110-1230220132212011-0213032313011013-3231133020120110-2121313230213320"></a>

### Direct properties for `items.get_spec`

- [infra](data-sources--site_registrations_by_state--reference--group-001.md#canonical-0233310023131030-0032203223303021-2221213101122020-0213011131203102-3002201310322332-2203103112111203-2313013132100313-0233320222120321): complete subsection reference.

- [passport](data-sources--site_registrations_by_state--reference--group-002.md#canonical-2203132300212312-1221232003303021-1330002313010203-1032320133313122-2323101201303111-2130030110023201-1121001200310221-1111100023322210): complete subsection reference.

<a id="canonical-1132231000102130-0300321223021303-0320311233021133-2332111303011102-0300003312323131-3213322113211203-3110021001322030-1201332000301100"></a>

<a id="canonical-0233301303000121-0322013222010030-1022112212023210-2200012202000320-1010130320023311-3233212110302300-3130130312011221-2131132131301112"></a>

#### `items.get_spec.token` property

Type: `"string"`. Computed.

Token is used for machine and tenant identification.

<a id="canonical-0233310023131030-0032203223303021-2221213101122020-0213011131203102-3002201310322332-2203103112111203-2313013132100313-0233320222120321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `items.get_spec.infra` properties

Breadcrumbs:

- [xcsh_site_registrations_by_state](../data-sources/site_registrations_by_state.md#canonical-1011310232230021-3303123312003111-3120023223111201-0320322111010320-3032200132321200-3310101021222131-0120032131300331-0201133223313303)
- [Property reference](data-sources--site_registrations_by_state--reference--group-001.md#canonical-0320011213231301-1102033233211220-0302131120233130-2122121213220203-2033211132132211-0333231313221113-3210300123023021-2021122221230130)
- [items](data-sources--site_registrations_by_state--reference--group-001.md#canonical-1133122031213013-2312233003121321-2003030111222220-0123212333122102-1113033120112122-1330232222330233-1111303223103223-2020000130121230)
- [items.get_spec](data-sources--site_registrations_by_state--reference--group-001.md#canonical-2321111233223323-1103131313132111-1332011333123121-2230221231113011-3312330012020320-0211132223121102-2102332200012230-0013300300311322)
- items.get_spec.infra

<a id="canonical-1002221311011210-1131112301131211-0320102012013000-1213223231322011-0320001222013330-3223110033211232-1322112232130020-1202223333330321"></a>

Type: `"single"`. Computed.

InfraMetadata stores information about instance infrastructure.

<a id="canonical-2313221102222122-2330101121330201-3131013302102233-0212131100103032-0233013121000113-3320100100332320-3213303321031002-2233002302231011"></a>

### Direct properties for `items.get_spec.infra`

<a id="canonical-2110001112031023-3001212111001230-2133311133312000-0023132113033323-3031020201203201-0123201101233103-0330000232130332-0200211112300223"></a>

#### `items.get_spec.infra.availability_zone` property

Type: `"string"`. Computed.

Availability Zone is a high-availability offering that protects your applications and data from
datacenter failures.

- [bond_config](data-sources--site_registrations_by_state--reference--group-001.md#canonical-1010031133200233-2032033102132113-0311123133113020-0330211122202233-1213021231220012-2111022032103301-0020131213033000-2013231120311120): complete subsection reference.

<a id="canonical-0111331000230202-1301310131133013-0033132310002300-0032133002031113-3001023200330302-0000102221112222-3121020212200011-0322222132032303"></a>

<a id="canonical-3332223333330233-0002110211212210-3212012031313210-1011232320331101-1302110223121231-1022302001002012-3211122312201131-3122212023310231"></a>

#### `items.get_spec.infra.certified_hw` property

Type: `"string"`. Computed.

Certified HW name used to map with F5XC certified\_hardware definition.

<a id="canonical-2211131130133310-3102212321321232-2231310210103010-2310311110032213-1010321031222331-1323023133221111-2200333030212000-3000023322220232"></a>

<a id="canonical-2103003033013001-2101330222322122-0112302301313130-3311222121313130-2230211003230333-0021221122300220-2122100133132011-0121011222310112"></a>

#### `items.get_spec.infra.domain` property

Type: `"string"`. Computed.

Machine domain. It's used for Kubernetes cloud provider when domain must be different than F5
Distributed Cloud.

<a id="canonical-2320322331101001-1212122110302303-1232330203103010-2332111310331122-2000221010112232-0203230321023202-1333113121113201-3333323012210112"></a>

<a id="canonical-0012213131010110-1012302213232121-0103210032001333-2331032110231200-0200312112021000-1212330310011023-1310321031131302-1331012313332231"></a>

#### `items.get_spec.infra.hostname` property

Type: `"string"`. Computed.

Must be unique in entire cluster and same as OS settings. '.' (dots) are not allowed in hostname.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 1024),
  stringvalidator.RegexMatches(regexp.MustCompile(`^([a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?\.)*[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?$`),
    ""),
}
```

- [hugepages](data-sources--site_registrations_by_state--reference--group-001.md#canonical-3111011221013121-0123133022212310-2313022103210212-3030032230300132-0112100223320222-1230313101033023-0301303313232221-0123311123101223): complete subsection reference.

- [hw_info](data-sources--site_registrations_by_state--reference--group-001.md#canonical-0221000332223002-0003222301201320-0320022333112331-1222022120030000-2003302110120110-1231101313210231-0010010012113111-1302023332323232): complete subsection reference.

<a id="canonical-1100012210010132-3013000233132301-1220112030020021-0223223211311330-0331033331312103-1103030230312212-3233203012320111-2030223022321113"></a>

<a id="canonical-1201112032302100-2220013202133001-3323033030013203-1312032212103230-2101312300021222-0131011332033330-2000012223011010-3210121331100133"></a>

#### `items.get_spec.infra.instance_id` property

Type: `"string"`. Computed.

Instance ID (assigned by infrastructure provider).

- [interfaces](data-sources--site_registrations_by_state--reference--group-002.md#canonical-3230330112200222-0121231210202301-2300313103221122-0122100021022000-1303123321310111-2000001221200201-0320020011130332-2230101100000000): complete subsection reference.

- [internet_proxy](data-sources--site_registrations_by_state--reference--group-002.md#canonical-3123211222220010-3010203323202310-0310321132111233-2302131232321211-3310321023331323-1320100102101121-3030032120322001-1122300330012121): complete subsection reference.

<a id="canonical-0221312133302033-0312322330000112-2212223133320221-3331231311012001-1131310113001113-2022232313231013-1311111332231201-1223331223130101"></a>

<a id="canonical-2311021212020200-1033133002120230-3323130133210303-2220211022121312-3311301031123121-1032303313230013-1320012123031220-3210020231202123"></a>

#### `items.get_spec.infra.is_slo_static` property

Type: `"bool"`. Computed.

Is SLO Static. Indicates whether the SLO is static.

<a id="canonical-2312101011230113-3012011001232032-2202203322201111-3311210032003121-3332000020033111-0312111213020221-1003212102132322-1323233013003003"></a>

<a id="canonical-2333022212031221-0221212111112203-3210313111013313-3313322031333330-0302322013003322-2221300231020321-2030201121320232-1212023030002003"></a>

#### `items.get_spec.infra.machine_id` property

Type: `"string"`. Computed.

Machine ID - generated by operating system.

<a id="canonical-0211022311130320-1332130001013222-0013030100303302-1222303121212333-3022323323210110-1201020002133323-0001103302211210-3232221133102123"></a>

<a id="canonical-2312021103021330-2023332323320010-2012233031332012-3333202011320333-0010203101330001-1233211103332122-1332212132001200-1331333121302312"></a>

#### `items.get_spec.infra.provider_ref` property

Type: `"string"`. Computed.

\[Enum:
UNKNOWN|AWS|GOOGLE|AZURE|VMWARE|KVM|OTHER|VOLTERRA|IBMCLOUD|UNKNOWN\_K8S|AWS\_K8S|GCP\_K8S|AZURE\_K8S|VMWARE\_K8S|KVM\_K8S|OTHER\_K8S|VOLTERRA\_K8S|IBMCLOUD\_K8S|F5OS|RSERIES|OCI|NUTANIX|OPENSTACK|EQUINIX|OPENSHIFT\_VIRTUALIZATION|KUBERNETES\]
Infrastructure provider enum for registration. It describes where is instance running. Provider was
not detected AWS cloud instance Google cloud instance Azure cloud instance VMWare VM KVM VM Other
provider, which was not identified by system. Possible values are \`UNKNOWN\`, \`AWS\`, \`GOOGLE\`,
\`AZURE\`, \`VMWARE\`, \`KVM\`, \`OTHER\`, \`VOLTERRA\`, \`IBMCLOUD\`, \`UNKNOWN\_K8S\`,
\`AWS\_K8S\`, \`GCP\_K8S\`, \`AZURE\_K8S\`, \`VMWARE\_K8S\`, \`KVM\_K8S\`, \`OTHER\_K8S\`,
\`VOLTERRA\_K8S\`, \`IBMCLOUD\_K8S\`, \`F5OS\`, \`RSERIES\`, \`OCI\`, \`NUTANIX\`, \`OPENSTACK\`,
\`EQUINIX\`, \`OPENSHIFT\_VIRTUALIZATION\`, \`KUBERNETES\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["AWS","AWS_K8S","AZURE","AZURE_K8S","EQUINIX","F5OS","GCP_K8S","GOOGLE","IBMCLOUD","IBMCLOUD_K8S","KUBERNETES","KVM","KVM_K8S","NUTANIX","OCI","OPENSHIFT_VIRTUALIZATION","OPENSTACK","OTHER","OTHER_K8S","RSERIES","UNKNOWN","UNKNOWN_K8S","VMWARE","VMWARE_K8S","VOLTERRA","VOLTERRA_K8S"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("UNKNOWN",
    "AWS",
    "GOOGLE",
    "AZURE",
    "VMWARE",
    "KVM",
    "OTHER",
    "VOLTERRA",
    "IBMCLOUD",
    "UNKNOWN_K8S",
    "AWS_K8S",
    "GCP_K8S",
    "AZURE_K8S",
    "VMWARE_K8S",
    "KVM_K8S",
    "OTHER_K8S",
    "VOLTERRA_K8S",
    "IBMCLOUD_K8S",
    "F5OS",
    "RSERIES",
    "OCI",
    "NUTANIX",
    "OPENSTACK",
    "EQUINIX",
    "OPENSHIFT_VIRTUALIZATION",
    "KUBERNETES"),
}
```

- [sw_info](data-sources--site_registrations_by_state--reference--group-002.md#canonical-0321201132012310-2302210020130033-2010303311202133-2130012331323002-1023122002033303-2213331203023221-2020123333223033-0031113322201121): complete subsection reference.

<a id="canonical-3233110202132322-2230223011110330-2020312101112131-1303000302023133-0301112301100120-1000131112001111-3001200132312223-2002101010202221"></a>

<a id="canonical-0122211001302121-2201130011333100-2110321332122131-1010211212001333-1201212332030221-3302030201201331-0120233233201130-2100212011121313"></a>

#### `items.get_spec.infra.timestamp` property

Type: `"string"`. Computed.

It's used to verify machine have acceptable time difference from server.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthBetween(20, 1024),
  stringvalidator.RegexMatches(regexp.MustCompile(`^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d{3})?Z?$`),
    ""),
}
```

<a id="canonical-3122120203201200-2012113013321233-1022001330220230-3211230102132131-0210223221133031-1011232323200313-0321330030112300-2330300321303033"></a>

<a id="canonical-3111102131311120-3011010100231311-3332122020131230-2221200313123233-2001210221203303-2022221303201232-2102032221230113-0012231321332212"></a>

#### `items.get_spec.infra.zone` property

Type: `"string"`. Computed.

Instance zone (or region), depends on provider.

<a id="canonical-1010031133200233-2032033102132113-0311123133113020-0330211122202233-1213021231220012-2111022032103301-0020131213033000-2013231120311120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `items.get_spec.infra.bond_config` properties

Breadcrumbs:

- [xcsh_site_registrations_by_state](../data-sources/site_registrations_by_state.md#canonical-1011310232230021-3303123312003111-3120023223111201-0320322111010320-3032200132321200-3310101021222131-0120032131300331-0201133223313303)
- [Property reference](data-sources--site_registrations_by_state--reference--group-001.md#canonical-0320011213231301-1102033233211220-0302131120233130-2122121213220203-2033211132132211-0333231313221113-3210300123023021-2021122221230130)
- [items](data-sources--site_registrations_by_state--reference--group-001.md#canonical-1133122031213013-2312233003121321-2003030111222220-0123212333122102-1113033120112122-1330232222330233-1111303223103223-2020000130121230)
- [items.get_spec](data-sources--site_registrations_by_state--reference--group-001.md#canonical-2321111233223323-1103131313132111-1332011333123121-2230221231113011-3312330012020320-0211132223121102-2102332200012230-0013300300311322)
- [items.get_spec.infra](data-sources--site_registrations_by_state--reference--group-001.md#canonical-0233310023131030-0032203223303021-2221213101122020-0213011131203102-3002201310322332-2203103112111203-2313013132100313-0233320222120321)
- items.get_spec.infra.bond_config

<a id="canonical-3303001302112112-0231233102030210-1231112300212210-2200030210011223-1331122120203131-2022210013213311-1020233332021000-2333323021311312"></a>

Type: `"single"`. Computed.

Bond device configuration for VPM registration.

<a id="canonical-2012311333012103-3333021000323322-2110211212010101-2102133300312300-3301103210110330-2023112231112213-1111121132321100-0221301203322031"></a>

### Direct properties for `items.get_spec.infra.bond_config`

<a id="canonical-3220330201100321-1112200101303010-0320313221220132-3303322201111320-0000012110313211-3330230023231301-1302132123202323-2011010130213002"></a>

#### `items.get_spec.infra.bond_config.interfaces` property

Type: `["list", "string"]`. Computed.

Member Interfaces. Configuration parameter for interfaces

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeBetween(1, 8),
}
```

<a id="canonical-1311111113310010-0021013003213320-2132131330012123-2303033221221331-0231131003111202-0231003232311001-3003220022100223-2102020331313323"></a>

<a id="canonical-3110111010210033-3223230313301031-3310103122322120-0011301231030133-1121021320310201-0101331123210301-2020003222310111-1120102133230001"></a>

#### `items.get_spec.infra.bond_config.mode` property

Type: `"string"`. Computed.

\[Enum: BOND\_MODE\_UNSPECIFIED|ACTIVE\_BACKUP|LACP\_802\_3AD\] Bonding mode for bond device
configuration Bond mode is not specified Active-backup bond mode (one interface active, others as
backup) IEEE 802.3ad Dynamic link aggregation (LACP). Possible values are
\`BOND\_MODE\_UNSPECIFIED\`, \`ACTIVE\_BACKUP\`, \`LACP\_802\_3AD\`. Defaults to
\`BOND\_MODE\_UNSPECIFIED\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["ACTIVE_BACKUP","BOND_MODE_UNSPECIFIED","LACP_802_3AD"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("BOND_MODE_UNSPECIFIED",
    "ACTIVE_BACKUP",
    "LACP_802_3AD"),
}
```

<a id="canonical-1122123313203213-0230032233323111-2120330221002100-2023010103010010-2130133211331120-3210310133331302-0030210113120323-3101310013112021"></a>

<a id="canonical-3112131102020113-3102131300020132-3011022003033310-1202333300123333-0200311313323120-1221213023202111-0121200311001212-0031232023202222"></a>

#### `items.get_spec.infra.bond_config.name` property

Type: `"string"`. Computed.

Bond Name. Human-readable name for the resource

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 64),
}
```

<a id="canonical-3111011221013121-0123133022212310-2313022103210212-3030032230300132-0112100223320222-1230313101033023-0301303313232221-0123311123101223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `items.get_spec.infra.hugepages` properties

Breadcrumbs:

- [xcsh_site_registrations_by_state](../data-sources/site_registrations_by_state.md#canonical-1011310232230021-3303123312003111-3120023223111201-0320322111010320-3032200132321200-3310101021222131-0120032131300331-0201133223313303)
- [Property reference](data-sources--site_registrations_by_state--reference--group-001.md#canonical-0320011213231301-1102033233211220-0302131120233130-2122121213220203-2033211132132211-0333231313221113-3210300123023021-2021122221230130)
- [items](data-sources--site_registrations_by_state--reference--group-001.md#canonical-1133122031213013-2312233003121321-2003030111222220-0123212333122102-1113033120112122-1330232222330233-1111303223103223-2020000130121230)
- [items.get_spec](data-sources--site_registrations_by_state--reference--group-001.md#canonical-2321111233223323-1103131313132111-1332011333123121-2230221231113011-3312330012020320-0211132223121102-2102332200012230-0013300300311322)
- [items.get_spec.infra](data-sources--site_registrations_by_state--reference--group-001.md#canonical-0233310023131030-0032203223303021-2221213101122020-0213011131203102-3002201310322332-2203103112111203-2313013132100313-0233320222120321)
- items.get_spec.infra.hugepages

<a id="canonical-3032120021203310-1331012033230302-3230030332013003-0023102002322223-3100013303220313-2110012212332232-3010111310201210-1332121023203011"></a>

Type: `"list"`. Computed.

Hugepage settings for CE on K8s SMV2 site.

<a id="canonical-1030312000013232-1312212102201211-3110331020131310-3301312223112222-1102123003311311-2002331301311120-2122220010111302-3020103122013320"></a>

### Direct properties for `items.get_spec.infra.hugepages`

<a id="canonical-2301032013200331-0113201303121231-2101002220010021-2233030223302121-3310100111120110-2300002102211000-3123300002300112-1120011230032102"></a>

#### `items.get_spec.infra.hugepages.free` property

Type: `"number"`. Computed.

Free Hugepages. Total number of free hugepages present.

<a id="canonical-3110310011232000-3213300011320001-1002212113223031-2120032013311103-0300320122331031-3320003113200111-2002030022103330-0200023212222130"></a>

<a id="canonical-0122333121101310-1330221131312010-0112233302210111-1002203233112033-3131231111213002-2220101211030031-2322101031232111-2323200003123230"></a>

#### `items.get_spec.infra.hugepages.page_size` property

Type: `"number"`. Computed.

Hugepage Size. Size of each hugepage.

<a id="canonical-2110311003230300-3203311023002010-2303003113211330-3012213131333330-0221032101330033-2132030133000303-0103231033021211-1211120132323312"></a>

<a id="canonical-1132131200302231-2232213112123122-0210033111330233-0020012332312211-2030201032112033-0013033122131230-1003132320122023-0313203223111110"></a>

#### `items.get_spec.infra.hugepages.total` property

Type: `"number"`. Computed.

Total Hugepages. Total number of hugepages present.

<a id="canonical-0221000332223002-0003222301201320-0320022333112331-1222022120030000-2003302110120110-1231101313210231-0010010012113111-1302023332323232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `items.get_spec.infra.hw_info` properties

Breadcrumbs:

- [xcsh_site_registrations_by_state](../data-sources/site_registrations_by_state.md#canonical-1011310232230021-3303123312003111-3120023223111201-0320322111010320-3032200132321200-3310101021222131-0120032131300331-0201133223313303)
- [Property reference](data-sources--site_registrations_by_state--reference--group-001.md#canonical-0320011213231301-1102033233211220-0302131120233130-2122121213220203-2033211132132211-0333231313221113-3210300123023021-2021122221230130)
- [items](data-sources--site_registrations_by_state--reference--group-001.md#canonical-1133122031213013-2312233003121321-2003030111222220-0123212333122102-1113033120112122-1330232222330233-1111303223103223-2020000130121230)
- [items.get_spec](data-sources--site_registrations_by_state--reference--group-001.md#canonical-2321111233223323-1103131313132111-1332011333123121-2230221231113011-3312330012020320-0211132223121102-2102332200012230-0013300300311322)
- [items.get_spec.infra](data-sources--site_registrations_by_state--reference--group-001.md#canonical-0233310023131030-0032203223303021-2221213101122020-0213011131203102-3002201310322332-2203103112111203-2313013132100313-0233320222120321)
- items.get_spec.infra.hw_info

<a id="canonical-0333310021202312-0202220020312133-0002332132033111-2132022013221100-3202322313233011-1230230331231023-3321102113223220-2002021322132131"></a>

Type: `"single"`. Computed.

OsInfo holds information about host OS and HW.

<a id="canonical-2313011030210001-1021110111210012-1023233220100311-1031222310322322-0030011022212201-1330331310312320-2200331210020110-1122310231230303"></a>

### Direct properties for `items.get_spec.infra.hw_info`

- [bios](data-sources--site_registrations_by_state--reference--group-001.md#canonical-2033212002211330-0221113203212011-0331220330030210-0212110133111023-3231110302031231-0321011211322101-3331333322123310-2003010003111320): complete subsection reference.

- [board](data-sources--site_registrations_by_state--reference--group-001.md#canonical-0302231022210121-1112232311012020-1012233002003102-1012113322030333-2003332033311013-2212200221330110-0021312230200131-3102320022322223): complete subsection reference.

- [chassis](data-sources--site_registrations_by_state--reference--group-001.md#canonical-2010322021010012-0322303300203232-0200220212222113-1232321232010110-2210323310223202-2313220002300032-2303013122211331-1112322022122133): complete subsection reference.

- [CPU](data-sources--site_registrations_by_state--reference--group-001.md#canonical-0112300230322103-3202303121113023-1203020223232020-0122333211032210-0110032120003333-1030010031203003-3111310212201112-0013022132131030): complete subsection reference.

- [GPU](data-sources--site_registrations_by_state--reference--group-001.md#canonical-2003000320003221-3330311032231112-2020130312301002-1001320221120301-1221202232331033-1300233323021203-0002320002301213-2112321230112211): complete subsection reference.

- [kernel](data-sources--site_registrations_by_state--reference--group-001.md#canonical-0321112322120311-1031123112113320-2110132121130031-0301303201200131-1310103312023122-1221302001022130-3033302103122112-0331031303233002): complete subsection reference.

- [memory](data-sources--site_registrations_by_state--reference--group-001.md#canonical-0320021223023322-2211221200033132-0223111232231321-3100100311210120-3122221233211010-3023100313001032-3332012213031313-1302111032131300): complete subsection reference.

- [network](data-sources--site_registrations_by_state--reference--group-001.md#canonical-3012301313312232-1222121333031330-2223120332121101-3221111311222111-2001133120320212-0200033232310103-3022101110212030-0010022023223111): complete subsection reference.

<a id="canonical-2103222010122313-2101300233130033-3020113032232000-2122103113221200-2122133223212320-2111001201211110-1310113212020330-3022203112301123"></a>

<a id="canonical-0221321312123321-3031121232013230-1003030221133121-3000110320213311-1212332323302121-0013101211003001-1112323230232031-1033033012311331"></a>

#### `items.get_spec.infra.hw_info.numa_nodes` property

Type: `"number"`. Computed.

Non-uniform memory access (NUMA) nodes count.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.AtLeast(0),
}
```

- [os](data-sources--site_registrations_by_state--reference--group-001.md#canonical-0221010331001303-0211231031231031-3111020113121012-2031133130213031-0212201031013210-0313203213011133-3233121100202112-0301332113302302): complete subsection reference.

- [product](data-sources--site_registrations_by_state--reference--group-001.md#canonical-2112103311331133-0033002130121030-1233121233310203-3100323312013313-2321213001202221-1313121320130013-3300003103323023-2012010322113133): complete subsection reference.

- [storage](data-sources--site_registrations_by_state--reference--group-002.md#canonical-2121200210023133-3230202221331302-0330102230131100-1112231300032211-3313013102203030-3023101123020000-3013113003000002-3121000001131113): complete subsection reference.

- [usb](data-sources--site_registrations_by_state--reference--group-002.md#canonical-1012320201023020-0031212333000222-1120323223011023-3332111323133201-0311030130133132-2012321100132030-0131311331010212-2130113133123220): complete subsection reference.

<a id="canonical-2033212002211330-0221113203212011-0331220330030210-0212110133111023-3231110302031231-0321011211322101-3331333322123310-2003010003111320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `items.get_spec.infra.hw_info.bios` properties

Breadcrumbs:

- [xcsh_site_registrations_by_state](../data-sources/site_registrations_by_state.md#canonical-1011310232230021-3303123312003111-3120023223111201-0320322111010320-3032200132321200-3310101021222131-0120032131300331-0201133223313303)
- [Property reference](data-sources--site_registrations_by_state--reference--group-001.md#canonical-0320011213231301-1102033233211220-0302131120233130-2122121213220203-2033211132132211-0333231313221113-3210300123023021-2021122221230130)
- [items](data-sources--site_registrations_by_state--reference--group-001.md#canonical-1133122031213013-2312233003121321-2003030111222220-0123212333122102-1113033120112122-1330232222330233-1111303223103223-2020000130121230)
- [items.get_spec](data-sources--site_registrations_by_state--reference--group-001.md#canonical-2321111233223323-1103131313132111-1332011333123121-2230221231113011-3312330012020320-0211132223121102-2102332200012230-0013300300311322)
- [items.get_spec.infra](data-sources--site_registrations_by_state--reference--group-001.md#canonical-0233310023131030-0032203223303021-2221213101122020-0213011131203102-3002201310322332-2203103112111203-2313013132100313-0233320222120321)
- [items.get_spec.infra.hw_info](data-sources--site_registrations_by_state--reference--group-001.md#canonical-0221000332223002-0003222301201320-0320022333112331-1222022120030000-2003302110120110-1231101313210231-0010010012113111-1302023332323232)
- items.get_spec.infra.hw_info.bios

<a id="canonical-0103130310322220-3032113002302323-0002030023120003-3020233013000012-0030001030011331-3301210013031021-1111201323221222-0202101330020313"></a>

Type: `"single"`. Computed.

Bios Data. BIOS information.

<a id="canonical-3323332123000203-1300023330222330-1001023013102213-3113311330230301-0033111200220310-3223032110200033-3213200010133331-3121030032230320"></a>

### Direct properties for `items.get_spec.infra.hw_info.bios`

<a id="canonical-3103113300332023-2122213333000301-3000333211200222-1322212332330022-0031203110202213-2111322221331121-2303001113021210-2013303020110320"></a>

#### `items.get_spec.infra.hw_info.bios.date` property

Type: `"string"`. Computed.

Information from /sys/class/dmi/ID/bios\_date.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthBetween(10, 1024),
  stringvalidator.RegexMatches(regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`),
    ""),
}
```

<a id="canonical-1202032230213123-2023202330222201-3100311103022303-0002112211301232-2002111222003110-2023300211102232-0303213333320122-2001331321121003"></a>

<a id="canonical-2030222322313330-0132022100302310-3200032121013231-3010133222101310-3202303222330102-2101332203123031-1320210033300220-0312223222110121"></a>

#### `items.get_spec.infra.hw_info.bios.vendor` property

Type: `"string"`. Computed.

Information from /sys/class/dmi/ID/bios\_vendor.

<a id="canonical-1331201120302122-0011320110112111-3013102313211331-2010202111220112-1000311112311123-1210303122000122-0212320110103013-2020002210133331"></a>

<a id="canonical-2101312122012130-3012113003221323-3303123003222012-2132301001021223-0111312031023230-3003311312300002-2332211130320120-0313311131100230"></a>

#### `items.get_spec.infra.hw_info.bios.version` property

Type: `"string"`. Computed.

Information from /sys/class/dmi/ID/bios\_version.

<a id="canonical-0302231022210121-1112232311012020-1012233002003102-1012113322030333-2003332033311013-2212200221330110-0021312230200131-3102320022322223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `items.get_spec.infra.hw_info.board` properties

Breadcrumbs:

- [xcsh_site_registrations_by_state](../data-sources/site_registrations_by_state.md#canonical-1011310232230021-3303123312003111-3120023223111201-0320322111010320-3032200132321200-3310101021222131-0120032131300331-0201133223313303)
- [Property reference](data-sources--site_registrations_by_state--reference--group-001.md#canonical-0320011213231301-1102033233211220-0302131120233130-2122121213220203-2033211132132211-0333231313221113-3210300123023021-2021122221230130)
- [items](data-sources--site_registrations_by_state--reference--group-001.md#canonical-1133122031213013-2312233003121321-2003030111222220-0123212333122102-1113033120112122-1330232222330233-1111303223103223-2020000130121230)
- [items.get_spec](data-sources--site_registrations_by_state--reference--group-001.md#canonical-2321111233223323-1103131313132111-1332011333123121-2230221231113011-3312330012020320-0211132223121102-2102332200012230-0013300300311322)
- [items.get_spec.infra](data-sources--site_registrations_by_state--reference--group-001.md#canonical-0233310023131030-0032203223303021-2221213101122020-0213011131203102-3002201310322332-2203103112111203-2313013132100313-0233320222120321)
- [items.get_spec.infra.hw_info](data-sources--site_registrations_by_state--reference--group-001.md#canonical-0221000332223002-0003222301201320-0320022333112331-1222022120030000-2003302110120110-1231101313210231-0010010012113111-1302023332323232)
- items.get_spec.infra.hw_info.board

<a id="canonical-0213010113120210-3032120313301031-3321220112022032-1212323310131213-0131033002013130-2130102013131130-0000103132020110-3023023022102022"></a>

Type: `"single"`. Computed.

Board Details. Board information.

<a id="canonical-0213202330023113-0020212331232322-2032012200112321-3311123100322221-2131310210320000-3220223031001012-1322322032031030-1033321330001103"></a>

### Direct properties for `items.get_spec.infra.hw_info.board`

<a id="canonical-2111132221130330-1112211010130323-3221221133003101-0031200112321212-1300221121101302-1033303023223223-0231330012331011-2213011101211321"></a>

#### `items.get_spec.infra.hw_info.board.asset_tag` property

Type: `"string"`. Computed.

Information from /sys/class/dmi/ID/board\_asset\_tag.

<a id="canonical-0032301103220210-0302131300203332-1003023012102023-3102113023101131-3311230022330202-0002013213111320-3220121103102020-2012211212331131"></a>

<a id="canonical-2100122121033022-1132201333032213-2013102122323223-2023122030230230-2203311110322122-0100132120323210-0102303022211033-0331002102200312"></a>

#### `items.get_spec.infra.hw_info.board.name` property

Type: `"string"`. Computed.

Information from /sys/class/dmi/ID/board\_name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

<a id="canonical-1212013112030130-1012112112212321-3102320321232203-2320001323303032-0113133213100322-1333222030311232-0030002023301212-0312232121320331"></a>

<a id="canonical-0133232203230130-2130012013331301-1222030110302111-2302232321102333-1033222320223301-3123100011303232-0210023023013300-3301300232221101"></a>

#### `items.get_spec.infra.hw_info.board.serial` property

Type: `"string"`. Computed.

Information from /sys/class/dmi/ID/board\_serial.

<a id="canonical-0113121021330000-0110301002101330-1213131003020230-0313331201303132-0331323130000010-0212200322303331-0321211230132320-2322120131102030"></a>

<a id="canonical-1322231320233211-1222030010101112-3221022111001211-3010203020231322-0110323130203212-0301103003211222-2213213133131201-0131330123312311"></a>

#### `items.get_spec.infra.hw_info.board.vendor` property

Type: `"string"`. Computed.

Information from /sys/class/dmi/ID/board\_vendor.

<a id="canonical-1220323231200000-3120100132230330-2133003110031212-0013322210133333-1023102030300113-0002120133300023-2231100032311200-0021131111123220"></a>

<a id="canonical-0120332131231220-2301001021111133-0121103301300220-0011222311232122-0330000023321013-1312000001220300-3103030203301033-3320302122302013"></a>

#### `items.get_spec.infra.hw_info.board.version` property

Type: `"string"`. Computed.

Information from /sys/class/dmi/ID/board\_version.

<a id="canonical-2010322021010012-0322303300203232-0200220212222113-1232321232010110-2210323310223202-2313220002300032-2303013122211331-1112322022122133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `items.get_spec.infra.hw_info.chassis` properties

Breadcrumbs:

- [xcsh_site_registrations_by_state](../data-sources/site_registrations_by_state.md#canonical-1011310232230021-3303123312003111-3120023223111201-0320322111010320-3032200132321200-3310101021222131-0120032131300331-0201133223313303)
- [Property reference](data-sources--site_registrations_by_state--reference--group-001.md#canonical-0320011213231301-1102033233211220-0302131120233130-2122121213220203-2033211132132211-0333231313221113-3210300123023021-2021122221230130)
- [items](data-sources--site_registrations_by_state--reference--group-001.md#canonical-1133122031213013-2312233003121321-2003030111222220-0123212333122102-1113033120112122-1330232222330233-1111303223103223-2020000130121230)
- [items.get_spec](data-sources--site_registrations_by_state--reference--group-001.md#canonical-2321111233223323-1103131313132111-1332011333123121-2230221231113011-3312330012020320-0211132223121102-2102332200012230-0013300300311322)
- [items.get_spec.infra](data-sources--site_registrations_by_state--reference--group-001.md#canonical-0233310023131030-0032203223303021-2221213101122020-0213011131203102-3002201310322332-2203103112111203-2313013132100313-0233320222120321)
- [items.get_spec.infra.hw_info](data-sources--site_registrations_by_state--reference--group-001.md#canonical-0221000332223002-0003222301201320-0320022333112331-1222022120030000-2003302110120110-1231101313210231-0010010012113111-1302023332323232)
- items.get_spec.infra.hw_info.chassis

<a id="canonical-3331100232201021-2300001011220320-1210013021023223-3020312101023133-1032213230222200-3031220200020021-2233231203112110-3213110330213213"></a>

Type: `"single"`. Computed.

Chassis Details. Chassis information.

<a id="canonical-1300322022133321-2030331122223213-0030213133122002-1220120322331003-1103231012003103-1302110233210233-0031100201331113-2232102131101001"></a>

### Direct properties for `items.get_spec.infra.hw_info.chassis`

<a id="canonical-2210300013210200-2031013103203131-3312203220210231-0233332030002201-2320310012302211-2121113321322123-0211003233022332-3302223021202332"></a>

#### `items.get_spec.infra.hw_info.chassis.asset_tag` property

Type: `"string"`. Computed.

Information from /sys/class/dmi/ID/chassis\_asset\_tag.

<a id="canonical-2132122103200322-1322300230100122-0223332311201222-0133221122221130-1301002221131311-2313310313303012-0213210323122313-0032020032321321"></a>

<a id="canonical-0220310110030032-0003201202303010-0110102001313202-1210232300222203-2100012033300201-1131110330210203-1110023023322312-0122201233123322"></a>

#### `items.get_spec.infra.hw_info.chassis.serial` property

Type: `"string"`. Computed.

Information from /sys/class/dmi/ID/chassis\_serial.

<a id="canonical-2321300112102020-0310302231223313-3111303220231002-2013102202332033-1131232233230130-1102111300023012-0310000132210030-1300210013203122"></a>

<a id="canonical-2132001301322333-3232222222030010-2122020102030001-1321230323133001-1110221311300313-1220011012331331-2321323023020331-3033201121003332"></a>

#### `items.get_spec.infra.hw_info.chassis.type` property

Type: `"number"`. Computed.

Information from /sys/class/dmi/ID/chassis\_type.

<a id="canonical-1221120330122010-2021130130033131-3323023332330322-1002313210301331-0222223032032013-0201103233022202-3322333123130322-1013210311213003"></a>

<a id="canonical-1002222030010212-0131201100200330-1321020031310211-1312213132122102-3210121021231130-1233222331020012-0323122123111012-1232032120223123"></a>

#### `items.get_spec.infra.hw_info.chassis.vendor` property

Type: `"string"`. Computed.

Information from /sys/class/dmi/ID/chassis\_vendor.

<a id="canonical-3330133000110023-2220001212322331-3313003310033312-1231000131320210-0111123130201032-0211123102011212-1031000000220331-3013021100103300"></a>

<a id="canonical-0123133212221010-3321001333102131-1212133201222113-1011000122003331-0222100100103020-3213233313300322-0103321330003012-3130032230102113"></a>

#### `items.get_spec.infra.hw_info.chassis.version` property

Type: `"string"`. Computed.

Information from /sys/class/dmi/ID/chassis\_version.

<a id="canonical-0112300230322103-3202303121113023-1203020223232020-0122333211032210-0110032120003333-1030010031203003-3111310212201112-0013022132131030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `items.get_spec.infra.hw_info.cpu` properties

Breadcrumbs:

- [xcsh_site_registrations_by_state](../data-sources/site_registrations_by_state.md#canonical-1011310232230021-3303123312003111-3120023223111201-0320322111010320-3032200132321200-3310101021222131-0120032131300331-0201133223313303)
- [Property reference](data-sources--site_registrations_by_state--reference--group-001.md#canonical-0320011213231301-1102033233211220-0302131120233130-2122121213220203-2033211132132211-0333231313221113-3210300123023021-2021122221230130)
- [items](data-sources--site_registrations_by_state--reference--group-001.md#canonical-1133122031213013-2312233003121321-2003030111222220-0123212333122102-1113033120112122-1330232222330233-1111303223103223-2020000130121230)
- [items.get_spec](data-sources--site_registrations_by_state--reference--group-001.md#canonical-2321111233223323-1103131313132111-1332011333123121-2230221231113011-3312330012020320-0211132223121102-2102332200012230-0013300300311322)
- [items.get_spec.infra](data-sources--site_registrations_by_state--reference--group-001.md#canonical-0233310023131030-0032203223303021-2221213101122020-0213011131203102-3002201310322332-2203103112111203-2313013132100313-0233320222120321)
- [items.get_spec.infra.hw_info](data-sources--site_registrations_by_state--reference--group-001.md#canonical-0221000332223002-0003222301201320-0320022333112331-1222022120030000-2003302110120110-1231101313210231-0010010012113111-1302023332323232)
- items.get_spec.infra.hw_info.CPU

<a id="canonical-0310321211201110-1303132132113220-1211322301212103-0311023102300023-0301213211202003-3131131103023131-1213330311023200-3201123311121111"></a>

Type: `"single"`. Computed.

CPU Information. CPU information.

<a id="canonical-0222223302201212-2011311010323130-2230213112030013-3300031331232131-1131232321010333-0333320321033101-1110210031133002-1031222032330210"></a>

### Direct properties for `items.get_spec.infra.hw_info.cpu`

<a id="canonical-0230022223223110-1130110131323312-2120012033311313-0321330110310323-1132303031202311-0012111133110133-2332201203021231-2102331122333221"></a>

#### `items.get_spec.infra.hw_info.cpu.cache` property

Type: `"number"`. Computed.

Cache. CPU cache size in KB.

<a id="canonical-2213203213303212-1012003030200213-0131211233032203-0001130123021103-0123303331322221-1303012220333110-0201331130132303-1303330130110333"></a>

<a id="canonical-1301233100110331-0012011200231012-1023321030103300-1100313311333232-3123013023113022-0222002111311103-2000313021130013-3011032113331122"></a>

#### `items.get_spec.infra.hw_info.cpu.cores` property

Type: `"number"`. Computed.

Cores. Number of physical CPU cores.

<a id="canonical-2031032103300221-0002103032033222-3230202001213102-0032002212033010-1112130101313232-2031203000100323-1110323033123322-1223020211311101"></a>

<a id="canonical-0013233333000230-1321002130000200-3012323123202032-2101123003210120-0030232210021121-2322201231313230-2213132012032232-3330211111131220"></a>

#### `items.get_spec.infra.hw_info.cpu.cpus` property

Type: `"number"`. Computed.

CPUs. Number of physical CPUs.

<a id="canonical-3233323222101212-3131023323202311-0202102030103323-1310003330013333-3230203332000200-0301031130223302-1001131322130220-0233030320202011"></a>

<a id="canonical-2112111312212032-3300212001233103-1230303223121031-3120120033133120-0032023332020033-3302212220020023-1123223130013100-1001002313001001"></a>

#### `items.get_spec.infra.hw_info.cpu.model` property

Type: `"string"`. Computed.

Model. CPU model

<a id="canonical-1101330301300003-2111102322231312-0032210213322221-2130101013232123-3023111203221230-2301313101313210-0221223202211203-3213121311302032"></a>

<a id="canonical-3120233320301323-2031233220220232-2110111231102302-1021212322230232-3100332011233132-1101312203320222-2113130212233031-0011323020221031"></a>

#### `items.get_spec.infra.hw_info.cpu.speed` property

Type: `"number"`. Computed.

Speed. CPU clock rate in MHz.

<a id="canonical-2302223310310221-0203223022102121-3212031231203212-1302110112303322-1221002032020010-2030103202032312-2303002132001231-3012212212133033"></a>

<a id="canonical-3122131311133012-1132030023202002-0122102223232331-2331133203121202-0022312210003303-2102310230301232-2132123320131103-1210030322200121"></a>

#### `items.get_spec.infra.hw_info.cpu.threads` property

Type: `"number"`. Computed.

Threads. Number of logical (HT) CPU cores.

<a id="canonical-1030322002120332-2102131112212311-1130221320202300-0110113303101210-3133221331122021-2032300300230113-3130123013303332-3222300000232330"></a>

<a id="canonical-3110033031121113-0200232330033001-0332010210013011-3112012031323032-3221132121131331-2122302103100022-3203203112301122-2220200022001330"></a>

#### `items.get_spec.infra.hw_info.cpu.vendor` property

Type: `"string"`. Computed.

Vendor. CPU vendor.

<a id="canonical-2003000320003221-3330311032231112-2020130312301002-1001320221120301-1221202232331033-1300233323021203-0002320002301213-2112321230112211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `items.get_spec.infra.hw_info.gpu` properties

Breadcrumbs:

- [xcsh_site_registrations_by_state](../data-sources/site_registrations_by_state.md#canonical-1011310232230021-3303123312003111-3120023223111201-0320322111010320-3032200132321200-3310101021222131-0120032131300331-0201133223313303)
- [Property reference](data-sources--site_registrations_by_state--reference--group-001.md#canonical-0320011213231301-1102033233211220-0302131120233130-2122121213220203-2033211132132211-0333231313221113-3210300123023021-2021122221230130)
- [items](data-sources--site_registrations_by_state--reference--group-001.md#canonical-1133122031213013-2312233003121321-2003030111222220-0123212333122102-1113033120112122-1330232222330233-1111303223103223-2020000130121230)
- [items.get_spec](data-sources--site_registrations_by_state--reference--group-001.md#canonical-2321111233223323-1103131313132111-1332011333123121-2230221231113011-3312330012020320-0211132223121102-2102332200012230-0013300300311322)
- [items.get_spec.infra](data-sources--site_registrations_by_state--reference--group-001.md#canonical-0233310023131030-0032203223303021-2221213101122020-0213011131203102-3002201310322332-2203103112111203-2313013132100313-0233320222120321)
- [items.get_spec.infra.hw_info](data-sources--site_registrations_by_state--reference--group-001.md#canonical-0221000332223002-0003222301201320-0320022333112331-1222022120030000-2003302110120110-1231101313210231-0010010012113111-1302023332323232)
- items.get_spec.infra.hw_info.GPU

<a id="canonical-1000211133003211-1233011031300312-3120201120011003-1001120121231011-1321220023230303-2020222323033220-0100133321320121-1233233312320302"></a>

Type: `"single"`. Computed.

GPU. GPU information on server.

<a id="canonical-3103230331203021-2201213020023322-0003003310030222-1213331030122322-0000000100131332-3322210212031111-2330330311110012-2123220220311022"></a>

### Direct properties for `items.get_spec.infra.hw_info.gpu`

<a id="canonical-1100120200032110-1313323001123003-3333301232011022-0112123312111023-3303100230011130-3301130023201022-3032312102301201-0221111032013332"></a>

#### `items.get_spec.infra.hw_info.gpu.cuda_version` property

Type: `"string"`. Computed.

Cuda Version. GPU Cuda Version.

<a id="canonical-1110110203130132-3103222010111220-3322220323210202-2121001223022102-3101003103231032-3303122112302330-3311113312023133-3223033203333020"></a>

<a id="canonical-1120333003132210-0203120323103033-2212122002301320-1002321222203323-1221303220201012-3123202321323312-3321221130023202-1122113320131311"></a>

#### `items.get_spec.infra.hw_info.gpu.driver_version` property

Type: `"string"`. Computed.

Driver Version. GPU Driver Version.

- [gpu_device](data-sources--site_registrations_by_state--reference--group-001.md#canonical-1233312223031333-3311132012133130-0330213321211120-2000200020120321-1021323323213202-1221312320320032-0030032103112321-1311032112001133): complete subsection reference.

<a id="canonical-1233312223031333-3311132012133130-0330213321211120-2000200020120321-1021323323213202-1221312320320032-0030032103112321-1311032112001133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `items.get_spec.infra.hw_info.gpu.gpu_device` properties

Breadcrumbs:

- [xcsh_site_registrations_by_state](../data-sources/site_registrations_by_state.md#canonical-1011310232230021-3303123312003111-3120023223111201-0320322111010320-3032200132321200-3310101021222131-0120032131300331-0201133223313303)
- [Property reference](data-sources--site_registrations_by_state--reference--group-001.md#canonical-0320011213231301-1102033233211220-0302131120233130-2122121213220203-2033211132132211-0333231313221113-3210300123023021-2021122221230130)
- [items](data-sources--site_registrations_by_state--reference--group-001.md#canonical-1133122031213013-2312233003121321-2003030111222220-0123212333122102-1113033120112122-1330232222330233-1111303223103223-2020000130121230)
- [items.get_spec](data-sources--site_registrations_by_state--reference--group-001.md#canonical-2321111233223323-1103131313132111-1332011333123121-2230221231113011-3312330012020320-0211132223121102-2102332200012230-0013300300311322)
- [items.get_spec.infra](data-sources--site_registrations_by_state--reference--group-001.md#canonical-0233310023131030-0032203223303021-2221213101122020-0213011131203102-3002201310322332-2203103112111203-2313013132100313-0233320222120321)
- [items.get_spec.infra.hw_info](data-sources--site_registrations_by_state--reference--group-001.md#canonical-0221000332223002-0003222301201320-0320022333112331-1222022120030000-2003302110120110-1231101313210231-0010010012113111-1302023332323232)
- [items.get_spec.infra.hw_info.gpu](data-sources--site_registrations_by_state--reference--group-001.md#canonical-2003000320003221-3330311032231112-2020130312301002-1001320221120301-1221202232331033-1300233323021203-0002320002301213-2112321230112211)
- items.get_spec.infra.hw_info.GPU.gpu_device

<a id="canonical-3130210131120230-1222222131120103-3232113110102031-3012211323322312-2233231322321122-2211333012333132-2201100130120202-2322312012323101"></a>

Type: `"list"`. Computed.

GPU devices. List of GPU devices in server.

<a id="canonical-2120321203112101-1233221030313133-1311211322110200-0012211213010101-0113110222333020-2331111111322202-3312211111301300-0103321103001313"></a>

### Direct properties for `items.get_spec.infra.hw_info.gpu.gpu_device`

<a id="canonical-3012122111010333-0331132220300313-1322103322103011-2300032212210131-1010212221111031-0002102321233221-0022013120300300-1120222100202221"></a>

#### `items.get_spec.infra.hw_info.gpu.gpu_device.id` property

Type: `"string"`. Computed.

GPU ID. GPU ID

<a id="canonical-1232301333120131-1133310211230020-0230223313023310-2313002112332203-3120101022211331-0301023223131211-0223323220320210-2203212010133220"></a>

<a id="canonical-3230312132323110-1002101023301202-2222013220330202-1330012321101333-1301302200323022-0303001303000133-2223102123223331-1103312202230233"></a>

#### `items.get_spec.infra.hw_info.gpu.gpu_device.processes` property

Type: `"string"`. Computed.

Processes. GPU Processes.

<a id="canonical-0102211020331210-0010203332211003-2011020200220030-0233003012232123-1320113031011103-2103201210012232-0330110321203120-0103021220302221"></a>

<a id="canonical-0300223320300013-2020101132010110-1311202002333322-1111310130021103-3133003232023202-3232102203322123-1230220030102103-0311313031003210"></a>

#### `items.get_spec.infra.hw_info.gpu.gpu_device.product_name` property

Type: `"string"`. Computed.

Product Name. GPU Product Name.

<a id="canonical-0321112322120311-1031123112113320-2110132121130031-0301303201200131-1310103312023122-1221302001022130-3033302103122112-0331031303233002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `items.get_spec.infra.hw_info.kernel` properties

Breadcrumbs:

- [xcsh_site_registrations_by_state](../data-sources/site_registrations_by_state.md#canonical-1011310232230021-3303123312003111-3120023223111201-0320322111010320-3032200132321200-3310101021222131-0120032131300331-0201133223313303)
- [Property reference](data-sources--site_registrations_by_state--reference--group-001.md#canonical-0320011213231301-1102033233211220-0302131120233130-2122121213220203-2033211132132211-0333231313221113-3210300123023021-2021122221230130)
- [items](data-sources--site_registrations_by_state--reference--group-001.md#canonical-1133122031213013-2312233003121321-2003030111222220-0123212333122102-1113033120112122-1330232222330233-1111303223103223-2020000130121230)
- [items.get_spec](data-sources--site_registrations_by_state--reference--group-001.md#canonical-2321111233223323-1103131313132111-1332011333123121-2230221231113011-3312330012020320-0211132223121102-2102332200012230-0013300300311322)
- [items.get_spec.infra](data-sources--site_registrations_by_state--reference--group-001.md#canonical-0233310023131030-0032203223303021-2221213101122020-0213011131203102-3002201310322332-2203103112111203-2313013132100313-0233320222120321)
- [items.get_spec.infra.hw_info](data-sources--site_registrations_by_state--reference--group-001.md#canonical-0221000332223002-0003222301201320-0320022333112331-1222022120030000-2003302110120110-1231101313210231-0010010012113111-1302023332323232)
- items.get_spec.infra.hw_info.kernel

<a id="canonical-3022210232033133-3022230012103200-3113331312220033-1131032331030111-1330123330100130-1231030122230020-1200213333221321-0030200233120202"></a>

Type: `"single"`. Computed.

Kernel. Kernel information.

<a id="canonical-1232201311002010-0012312012023032-3101001012210013-0313030111333322-3201123323310012-1031311110302021-1123023023121032-1023002213201130"></a>

### Direct properties for `items.get_spec.infra.hw_info.kernel`

<a id="canonical-2301332113310132-2121313133212023-0101132220102001-2233213220322123-3110203233100130-3131130010012311-1332312200123000-1021130023222033"></a>

#### `items.get_spec.infra.hw_info.kernel.architecture` property

Type: `"string"`. Computed.

Architecture. Kernel architecture.

<a id="canonical-1222210020132301-3210130023111201-1301323110202122-3100313112210002-3330010002203032-3021323123122322-1330023022110031-1103120023130200"></a>

<a id="canonical-1033103320223001-2231131131011211-0230201031310113-3320001113331001-0112331200123030-0002231100212033-1332223013323013-3111012011310303"></a>

#### `items.get_spec.infra.hw_info.kernel.release` property

Type: `"string"`. Computed.

Release. Kernel release.

<a id="canonical-2223332112123102-2302323111103133-2103331300301020-0323303001310022-1303010003131321-2310000121200111-1202102333333220-0103110013012332"></a>

<a id="canonical-0122132021302212-3221201312200322-3331020221312322-3320202013321300-1310030001201330-3202022032220021-1013032100213312-2230333200012330"></a>

#### `items.get_spec.infra.hw_info.kernel.version` property

Type: `"string"`. Computed.

Version. Kernel version.

<a id="canonical-0320021223023322-2211221200033132-0223111232231321-3100100311210120-3122221233211010-3023100313001032-3332012213031313-1302111032131300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `items.get_spec.infra.hw_info.memory` properties

Breadcrumbs:

- [xcsh_site_registrations_by_state](../data-sources/site_registrations_by_state.md#canonical-1011310232230021-3303123312003111-3120023223111201-0320322111010320-3032200132321200-3310101021222131-0120032131300331-0201133223313303)
- [Property reference](data-sources--site_registrations_by_state--reference--group-001.md#canonical-0320011213231301-1102033233211220-0302131120233130-2122121213220203-2033211132132211-0333231313221113-3210300123023021-2021122221230130)
- [items](data-sources--site_registrations_by_state--reference--group-001.md#canonical-1133122031213013-2312233003121321-2003030111222220-0123212333122102-1113033120112122-1330232222330233-1111303223103223-2020000130121230)
- [items.get_spec](data-sources--site_registrations_by_state--reference--group-001.md#canonical-2321111233223323-1103131313132111-1332011333123121-2230221231113011-3312330012020320-0211132223121102-2102332200012230-0013300300311322)
- [items.get_spec.infra](data-sources--site_registrations_by_state--reference--group-001.md#canonical-0233310023131030-0032203223303021-2221213101122020-0213011131203102-3002201310322332-2203103112111203-2313013132100313-0233320222120321)
- [items.get_spec.infra.hw_info](data-sources--site_registrations_by_state--reference--group-001.md#canonical-0221000332223002-0003222301201320-0320022333112331-1222022120030000-2003302110120110-1231101313210231-0010010012113111-1302023332323232)
- items.get_spec.infra.hw_info.memory

<a id="canonical-0130112022201013-1331113032132201-0003330032103203-3121102022232113-3321331223231213-0313331211101203-1230212112331112-0222200101213223"></a>

Type: `"single"`. Computed.

Memory Information. Memory information.

<a id="canonical-1302230110231131-2301013210111212-3222213310131120-2203323203221011-3000103012232000-0031122030113222-0033301001123101-3031310331310101"></a>

### Direct properties for `items.get_spec.infra.hw_info.memory`

<a id="canonical-1033131011102131-2213133302000101-3312001301220031-0133113110030120-0212303011103321-0000022030322012-0001233332101202-2032311001100111"></a>

#### `items.get_spec.infra.hw_info.memory.size_mb` property

Type: `"number"`. Computed.

RAM. RAM size in MB.

<a id="canonical-2130110212310321-3220300310113111-2331023030232001-3233103320322111-2311322301210110-2101211120303123-3213203220201130-0201203111322203"></a>

<a id="canonical-2021332031320310-2113113020200023-3313132120231013-1300002212312100-2102012332331221-1210011101131323-3022021131313203-2223302231300122"></a>

#### `items.get_spec.infra.hw_info.memory.speed` property

Type: `"number"`. Computed.

Speed. RAM data rate in MT/s.

<a id="canonical-2312100222112212-3230333033313330-3200002220223301-1120200221121333-1102012101323333-0122003210033003-1210133112121103-1110121231012323"></a>

<a id="canonical-3330212322232020-1112022220010021-1321032230121221-0100222320323303-3320023103130220-1221113330002233-0030200132202330-3002020220113133"></a>

#### `items.get_spec.infra.hw_info.memory.type` property

Type: `"string"`. Computed.

Type. Type of memory, eg. DDR4.

<a id="canonical-3012301313312232-1222121333031330-2223120332121101-3221111311222111-2001133120320212-0200033232310103-3022101110212030-0010022023223111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `items.get_spec.infra.hw_info.network` properties

Breadcrumbs:

- [xcsh_site_registrations_by_state](../data-sources/site_registrations_by_state.md#canonical-1011310232230021-3303123312003111-3120023223111201-0320322111010320-3032200132321200-3310101021222131-0120032131300331-0201133223313303)
- [Property reference](data-sources--site_registrations_by_state--reference--group-001.md#canonical-0320011213231301-1102033233211220-0302131120233130-2122121213220203-2033211132132211-0333231313221113-3210300123023021-2021122221230130)
- [items](data-sources--site_registrations_by_state--reference--group-001.md#canonical-1133122031213013-2312233003121321-2003030111222220-0123212333122102-1113033120112122-1330232222330233-1111303223103223-2020000130121230)
- [items.get_spec](data-sources--site_registrations_by_state--reference--group-001.md#canonical-2321111233223323-1103131313132111-1332011333123121-2230221231113011-3312330012020320-0211132223121102-2102332200012230-0013300300311322)
- [items.get_spec.infra](data-sources--site_registrations_by_state--reference--group-001.md#canonical-0233310023131030-0032203223303021-2221213101122020-0213011131203102-3002201310322332-2203103112111203-2313013132100313-0233320222120321)
- [items.get_spec.infra.hw_info](data-sources--site_registrations_by_state--reference--group-001.md#canonical-0221000332223002-0003222301201320-0320022333112331-1222022120030000-2003302110120110-1231101313210231-0010010012113111-1302023332323232)
- items.get_spec.infra.hw_info.network

<a id="canonical-3102231021111033-1103120300102101-2212002122332223-2013221212101312-2111222323023200-1323011133301301-1320003132130232-0213331101221223"></a>

Type: `"list"`. Computed.

Network. List of network devices in server.

<a id="canonical-3110100233111332-1233022231010120-2203302222222303-1221113001330222-1330131322022000-1020323112011223-0103130110012220-2012021302030302"></a>

### Direct properties for `items.get_spec.infra.hw_info.network`

<a id="canonical-1322021022001323-2023211233222233-2113232223030213-3322231030031322-1211200011010122-1031332022001000-1221200033012323-1321110323022123"></a>

#### `items.get_spec.infra.hw_info.network.driver` property

Type: `"string"`. Computed.

Driver. Driver of device, eg. E1000e.

<a id="canonical-2000001002333001-1033002330303333-2233220102122003-2301211300021111-3311011103010102-0113323232321322-2320003302320322-1111112202302231"></a>

<a id="canonical-2202201022110110-1321111333311022-3122031003103301-1003313223230230-0100102112133213-1321313000203033-3120033011333121-1033312203211210"></a>

#### `items.get_spec.infra.hw_info.network.ip_address` property

Type: `["list", "string"]`. Computed.

IP Address. IP address on interface.

<a id="canonical-2302331023020230-0103002311100330-2331203000003101-2120013103302013-1300231032031120-3023011202321031-3220013221331323-2203113102201102"></a>

<a id="canonical-0200003231002323-2111120322010333-3122100113121312-2220300011121021-2122111020010001-0022022321220312-1030110021033300-0110101332110033"></a>

#### `items.get_spec.infra.hw_info.network.link_quality` property

Type: `"string"`. Computed.

\[Enum: QUALITY\_UNKNOWN|QUALITY\_GOOD|QUALITY\_POOR|QUALITY\_DISABLED\] Link quality determined by
VER using different probes Unknown quality Link quality is good Link quality is poor Quality
disabled. Possible values are \`QUALITY\_UNKNOWN\`, \`QUALITY\_GOOD\`, \`QUALITY\_POOR\`,
\`QUALITY\_DISABLED\`. Defaults to \`QUALITY\_UNKNOWN\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["QUALITY_DISABLED","QUALITY_GOOD","QUALITY_POOR","QUALITY_UNKNOWN"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("QUALITY_UNKNOWN",
    "QUALITY_GOOD",
    "QUALITY_POOR",
    "QUALITY_DISABLED"),
}
```

<a id="canonical-2301332301233202-0130312323133003-3001021020312003-0211233212023022-1131110103031020-1223321231110210-3010312113132001-0032001022311321"></a>

<a id="canonical-3211023320333322-3000203123110103-1202301322130311-0233331313321003-0221020111132002-0310232232213223-1321003032230211-2233321002322321"></a>

#### `items.get_spec.infra.hw_info.network.link_type` property

Type: `"string"`. Computed.

\[Enum:
LINK\_TYPE\_UNKNOWN|LINK\_TYPE\_ETHERNET|LINK\_TYPE\_WIFI\_802\_11AC|LINK\_TYPE\_WIFI\_802\_11BGN|LINK\_TYPE\_4G|LINK\_TYPE\_WIFI|LINK\_TYPE\_WAN\]
Link type of interface determined operationally Link type unknown Link type ethernet Wi-Fi link of
type 802.11ac Wi-Fi link of type 802.11bgn Link type 4G Wi-Fi link Wan link. Possible values are
\`LINK\_TYPE\_UNKNOWN\`, \`LINK\_TYPE\_ETHERNET\`, \`LINK\_TYPE\_WIFI\_802\_11AC\`,
\`LINK\_TYPE\_WIFI\_802\_11BGN\`, \`LINK\_TYPE\_4G\`, \`LINK\_TYPE\_WIFI\`, \`LINK\_TYPE\_WAN\`.
Defaults to \`LINK\_TYPE\_UNKNOWN\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["LINK_TYPE_4G","LINK_TYPE_ETHERNET","LINK_TYPE_UNKNOWN","LINK_TYPE_WAN","LINK_TYPE_WIFI","LINK_TYPE_WIFI_802_11AC","LINK_TYPE_WIFI_802_11BGN"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("LINK_TYPE_UNKNOWN",
    "LINK_TYPE_ETHERNET",
    "LINK_TYPE_WIFI_802_11AC",
    "LINK_TYPE_WIFI_802_11BGN",
    "LINK_TYPE_4G",
    "LINK_TYPE_WIFI",
    "LINK_TYPE_WAN"),
}
```

<a id="canonical-3130032032203332-2330020313100232-2312010130211321-3322131032133023-2021113011322013-1020313303033133-3200030230322103-1130203013133112"></a>

<a id="canonical-0323120303030203-2132333311131123-0013212011322210-3003313030103013-0203331320202112-0211002132113010-2033331223132333-3000233232322223"></a>

#### `items.get_spec.infra.hw_info.network.mac_address` property

Type: `"string"`. Computed.

MAC Address. MAC address on interface.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthBetween(17, 1024),
  stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})$`),
    ""),
}
```

<a id="canonical-1110101333221200-0023201313200201-2302300012031010-0300030201111302-1011030232022302-1220112323231220-2300133002232323-0110132230130110"></a>

<a id="canonical-2303301202303023-2002012132223210-1100330222312210-1112111313000211-3131301202102303-0131110220013030-2230302123030033-1023202011010023"></a>

#### `items.get_spec.infra.hw_info.network.name` property

Type: `"string"`. Computed.

Name. Name of device, eg. Eth0.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

<a id="canonical-2110010322100311-1122331232002331-0322332212010121-0222031003000233-0023233030322120-1233131231100302-2201102002302111-0113111313101213"></a>

<a id="canonical-1101030020300223-2010200023021223-2212301310032023-3230020110021021-2003132103133000-2003133112331320-3301113333133201-3300031133302120"></a>

#### `items.get_spec.infra.hw_info.network.port` property

Type: `"string"`. Computed.

Port. Used port, eg. Tp.

<a id="canonical-2311111111303331-1210003131233313-2232032100231021-3032311032011202-1123012011100211-3312302020230233-3001123103222132-1310020302323203"></a>

<a id="canonical-1111110132002100-3312133121322011-0222111232311230-3312310131202320-3322301022202100-3210031103012302-2333323322030322-3001123132301102"></a>

#### `items.get_spec.infra.hw_info.network.speed` property

Type: `"number"`. Computed.

Speed. Device max supported speed in Mbps.

<a id="canonical-0221010331001303-0211231031231031-3111020113121012-2031133130213031-0212201031013210-0313203213011133-3233121100202112-0301332113302302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `items.get_spec.infra.hw_info.os` properties

Breadcrumbs:

- [xcsh_site_registrations_by_state](../data-sources/site_registrations_by_state.md#canonical-1011310232230021-3303123312003111-3120023223111201-0320322111010320-3032200132321200-3310101021222131-0120032131300331-0201133223313303)
- [Property reference](data-sources--site_registrations_by_state--reference--group-001.md#canonical-0320011213231301-1102033233211220-0302131120233130-2122121213220203-2033211132132211-0333231313221113-3210300123023021-2021122221230130)
- [items](data-sources--site_registrations_by_state--reference--group-001.md#canonical-1133122031213013-2312233003121321-2003030111222220-0123212333122102-1113033120112122-1330232222330233-1111303223103223-2020000130121230)
- [items.get_spec](data-sources--site_registrations_by_state--reference--group-001.md#canonical-2321111233223323-1103131313132111-1332011333123121-2230221231113011-3312330012020320-0211132223121102-2102332200012230-0013300300311322)
- [items.get_spec.infra](data-sources--site_registrations_by_state--reference--group-001.md#canonical-0233310023131030-0032203223303021-2221213101122020-0213011131203102-3002201310322332-2203103112111203-2313013132100313-0233320222120321)
- [items.get_spec.infra.hw_info](data-sources--site_registrations_by_state--reference--group-001.md#canonical-0221000332223002-0003222301201320-0320022333112331-1222022120030000-2003302110120110-1231101313210231-0010010012113111-1302023332323232)
- items.get_spec.infra.hw_info.os

<a id="canonical-3130030000102013-2332131113101102-0200331231331111-3112130120211101-3320303202220312-3002320231012122-2013311013022232-3230221001303231"></a>

Type: `"single"`. Computed.

OS. Details of Operating System.

<a id="canonical-2032200220002220-0201302100203321-2323200323203202-3202121030201033-3021210202333322-0110233102332111-3210210112133122-3133312303023220"></a>

### Direct properties for `items.get_spec.infra.hw_info.os`

<a id="canonical-2211230221112302-3010000203133202-3121202332002233-0303313101323000-0333102022101101-2203303030030123-0120031010110022-2022000111310312"></a>

#### `items.get_spec.infra.hw_info.os.architecture` property

Type: `"string"`. Computed.

Architecture. Architecture of OS.

<a id="canonical-1110231003230230-2212313002303202-1313012030133023-0320311333132333-3010102211123121-1132010222012302-0200111002013003-3003021022330231"></a>

<a id="canonical-3120011123031232-0232122123023302-1313032110310333-2110022032013100-1322112330000102-0212211332213230-2302121223012200-2200300113302113"></a>

#### `items.get_spec.infra.hw_info.os.name` property

Type: `"string"`. Computed.

Name. Name of OS.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

<a id="canonical-1002103113013123-2103101322031010-3330211102320232-0200301200222120-2030211332003233-0010312302022123-0022213221130032-3110323033132102"></a>

<a id="canonical-0220333133001033-1203321132223113-0011222031122113-3021201213231311-2330033323302322-3031211311220221-3302112010233303-1233023201011103"></a>

#### `items.get_spec.infra.hw_info.os.release` property

Type: `"string"`. Computed.

Release. Release of the OS.

<a id="canonical-0223310213023122-1300132030302221-0000020333233000-1200100222211022-3322003012210222-1000333122330103-3233233210220023-1002002330101331"></a>

<a id="canonical-2321220101013333-3203300221320311-0100013011022200-0020001222210221-1030022132123123-3011100013211120-3322100110010203-2213322121001201"></a>

#### `items.get_spec.infra.hw_info.os.vendor` property

Type: `"string"`. Computed.

Vendor. Vendor of OS.

<a id="canonical-2322312213113031-2033103210201303-1333210333320010-2003121321323102-2111122131120033-0121320233122223-2201000120330010-1122013320030130"></a>

<a id="canonical-0113103320001111-1313332330120203-0120201123203211-2321321012003321-2311301232120302-2301131023201222-2001031000131233-1221231311001123"></a>

#### `items.get_spec.infra.hw_info.os.version` property

Type: `"string"`. Computed.

Version. Version of OS.

<a id="canonical-2112103311331133-0033002130121030-1233121233310203-3100323312013313-2321213001202221-1313121320130013-3300003103323023-2012010322113133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `items.get_spec.infra.hw_info.product` properties

Breadcrumbs:

- [xcsh_site_registrations_by_state](../data-sources/site_registrations_by_state.md#canonical-1011310232230021-3303123312003111-3120023223111201-0320322111010320-3032200132321200-3310101021222131-0120032131300331-0201133223313303)
- [Property reference](data-sources--site_registrations_by_state--reference--group-001.md#canonical-0320011213231301-1102033233211220-0302131120233130-2122121213220203-2033211132132211-0333231313221113-3210300123023021-2021122221230130)
- [items](data-sources--site_registrations_by_state--reference--group-001.md#canonical-1133122031213013-2312233003121321-2003030111222220-0123212333122102-1113033120112122-1330232222330233-1111303223103223-2020000130121230)
- [items.get_spec](data-sources--site_registrations_by_state--reference--group-001.md#canonical-2321111233223323-1103131313132111-1332011333123121-2230221231113011-3312330012020320-0211132223121102-2102332200012230-0013300300311322)
- [items.get_spec.infra](data-sources--site_registrations_by_state--reference--group-001.md#canonical-0233310023131030-0032203223303021-2221213101122020-0213011131203102-3002201310322332-2203103112111203-2313013132100313-0233320222120321)
- [items.get_spec.infra.hw_info](data-sources--site_registrations_by_state--reference--group-001.md#canonical-0221000332223002-0003222301201320-0320022333112331-1222022120030000-2003302110120110-1231101313210231-0010010012113111-1302023332323232)
- items.get_spec.infra.hw_info.product

<a id="canonical-1223313123203113-3320111012322010-2302013000111230-3211022213332132-2333123033123122-2012020221310313-1021021323232202-1221110021131200"></a>

Type: `"single"`. Computed.

Product Information. Product information.

<a id="canonical-1130121330221001-2020012321211120-2103020211232031-1012311320303312-1212323212311303-1202020300010120-2322012033033203-3321222223233332"></a>

### Direct properties for `items.get_spec.infra.hw_info.product`

<a id="canonical-0312113101230321-1020300101000111-3110332232222030-2112311232320011-1303212120330102-1101023122233322-3201231202112023-1300333213102213"></a>

#### `items.get_spec.infra.hw_info.product.name` property

Type: `"string"`. Computed.

Name. Product name, eg. For AWS m5a.xlarge. Info taken from /sys/class/dmi/ID/product\_name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

<a id="canonical-3130330123221303-3211020003012013-1012132332030313-1322203233333201-0322103023130111-3001222120330001-1320211332312200-1200302121212020"></a>

<a id="canonical-2311313312333330-3113202223321321-1022223002101130-1220033111023301-1232222211030202-0012201122310020-2111100220322130-3010012021231113"></a>

#### `items.get_spec.infra.hw_info.product.serial` property

Type: `"string"`. Computed.

Serial number, eg. For AWS 00000000-0000-4000-8000-23460f645a1f. Info taken from
/sys/class/dmi/ID/product\_serial.

<a id="canonical-2121132001111222-1222121021103311-3233020022102332-2333231331100213-0212112232030322-1222001330330330-1203011122311102-3003320331021032"></a>

<a id="canonical-3300200010100323-2313311030231031-3330310002102320-2030200302311220-1200121021200313-3331231300302230-1013113232332032-2310310003010203"></a>

#### `items.get_spec.infra.hw_info.product.vendor` property

Type: `"string"`. Computed.

Vendor. Vendor name, eg. For AWS Amazon EC2. Info taken from /sys/class/dmi/ID/product\_vendor.

<a id="canonical-1020213222201123-3322323121110223-0122010302203201-3313131311301100-3123010331312330-0202010113131300-2233311312312312-0010001020202310"></a>

<a id="canonical-3101000132110010-1121323120021103-3202000213331111-2031211101110211-0112313220332132-1133202102331210-2212322223122132-2310231022130210"></a>

#### `items.get_spec.infra.hw_info.product.version` property

Type: `"string"`. Computed.

Version name. Info taken from /sys/class/dmi/ID/product\_version.
