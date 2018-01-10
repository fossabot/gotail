package tail

const (
	S_MAGIC_AAFS            = 0x5A3C69F0
	S_MAGIC_ACFS            = 0x61636673
	S_MAGIC_ADFS            = 0xADF5
	S_MAGIC_AFFS            = 0xADFF
	S_MAGIC_AFS             = 0x5346414F
	S_MAGIC_ANON_INODE_FS   = 0x09041934
	S_MAGIC_APFS            = 24
	S_MAGIC_AUFS            = 0x61756673
	S_MAGIC_AUTOFS          = 0x0187
	S_MAGIC_BALLOON_KVM     = 0x13661366
	S_MAGIC_BEFS            = 0x42465331
	S_MAGIC_BDEVFS          = 0x62646576
	S_MAGIC_BFS             = 0x1BADFACE
	S_MAGIC_BPF_FS          = 0xCAFE4A11
	S_MAGIC_BINFMTFS        = 0x42494E4D
	S_MAGIC_BTRFS           = 0x9123683E
	S_MAGIC_BTRFS_TEST      = 0x73727279
	S_MAGIC_CEPH            = 0x00C36400
	S_MAGIC_CGROUP          = 0x0027E0EB
	S_MAGIC_CGROUP2         = 0x63677270
	S_MAGIC_CIFS            = 0xFF534D42
	S_MAGIC_CODA            = 0x73757245
	S_MAGIC_COH             = 0x012FF7B7
	S_MAGIC_CONFIGFS        = 0x62656570
	S_MAGIC_CRAMFS          = 0x28CD3D45
	S_MAGIC_CRAMFS_WEND     = 0x453DCD28
	S_MAGIC_DAXFS           = 0x64646178
	S_MAGIC_DEBUGFS         = 0x64626720
	S_MAGIC_DEVFS           = 0x1373
	S_MAGIC_DEVPTS          = 0x1CD1
	S_MAGIC_ECRYPTFS        = 0xF15F
	S_MAGIC_EFIVARFS        = 0xDE5E81E4
	S_MAGIC_EFS             = 0x00414A53
	S_MAGIC_EXOFS           = 0x5DF5
	S_MAGIC_EXT             = 0x137D
	S_MAGIC_EXT2            = 0xEF53
	S_MAGIC_EXT2_OLD        = 0xEF51
	S_MAGIC_F2FS            = 0xF2F52010
	S_MAGIC_FAT             = 0x4006
	S_MAGIC_FHGFS           = 0x19830326
	S_MAGIC_FUSEBLK         = 0x65735546
	S_MAGIC_FUSECTL         = 0x65735543
	S_MAGIC_FUTEXFS         = 0x0BAD1DEA
	S_MAGIC_GFS             = 0x01161970
	S_MAGIC_GPFS            = 0x47504653
	S_MAGIC_HFS             = 0x4244
	S_MAGIC_HFS_PLUS        = 0x482B
	S_MAGIC_HFS_X           = 0x4858
	S_MAGIC_HOSTFS          = 0x00C0FFEE
	S_MAGIC_HPFS            = 0xF995E849
	S_MAGIC_HUGETLBFS       = 0x958458F6
	S_MAGIC_MTD_INODE_FS    = 0x11307854
	S_MAGIC_IBRIX           = 0x013111A8
	S_MAGIC_INOTIFYFS       = 0x2BAD1DEA
	S_MAGIC_ISOFS           = 0x9660
	S_MAGIC_ISOFS_R_WIN     = 0x4004
	S_MAGIC_ISOFS_WIN       = 0x4000
	S_MAGIC_JFFS            = 0x07C0
	S_MAGIC_JFFS2           = 0x72B6
	S_MAGIC_JFS             = 0x3153464A
	S_MAGIC_KAFS            = 0x6B414653
	S_MAGIC_LOGFS           = 0xC97E8168
	S_MAGIC_LUSTRE          = 0x0BD00BD0
	S_MAGIC_M1FS            = 0x5346314D
	S_MAGIC_MINIX           = 0x137F
	S_MAGIC_MINIX_30        = 0x138F
	S_MAGIC_MINIX_V2        = 0x2468
	S_MAGIC_MINIX_V2_30     = 0x2478
	S_MAGIC_MINIX_V3        = 0x4D5A
	S_MAGIC_MQUEUE          = 0x19800202
	S_MAGIC_MSDOS           = 0x4D44
	S_MAGIC_NCP             = 0x564C
	S_MAGIC_NFS             = 0x6969
	S_MAGIC_NFSD            = 0x6E667364
	S_MAGIC_NILFS           = 0x3434
	S_MAGIC_NSFS            = 0x6E736673
	S_MAGIC_NTFS            = 0x5346544E
	S_MAGIC_OPENPROM        = 0x9FA1
	S_MAGIC_OCFS2           = 0x7461636F
	S_MAGIC_OVERLAYFS       = 0x794C7630
	S_MAGIC_PANFS           = 0xAAD7AAEA
	S_MAGIC_PIPEFS          = 0x50495045
	S_MAGIC_PRL_FS          = 0x7C7C6673
	S_MAGIC_PROC            = 0x9FA0
	S_MAGIC_PSTOREFS        = 0x6165676C
	S_MAGIC_QNX4            = 0x002F
	S_MAGIC_QNX6            = 0x68191122
	S_MAGIC_RAMFS           = 0x858458F6
	S_MAGIC_RDTGROUP        = 0x07655821
	S_MAGIC_REISERFS        = 0x52654973
	S_MAGIC_ROMFS           = 0x7275
	S_MAGIC_RPC_PIPEFS      = 0x67596969
	S_MAGIC_SECURITYFS      = 0x73636673
	S_MAGIC_SELINUX         = 0xF97CFF8C
	S_MAGIC_SMACK           = 0x43415D53
	S_MAGIC_SMB             = 0x517B
	S_MAGIC_SMB2            = 0xFE534D42
	S_MAGIC_SNFS            = 0xBEEFDEAD
	S_MAGIC_SOCKFS          = 0x534F434B
	S_MAGIC_SQUASHFS        = 0x73717368
	S_MAGIC_SYSFS           = 0x62656572
	S_MAGIC_SYSV2           = 0x012FF7B6
	S_MAGIC_SYSV4           = 0x012FF7B5
	S_MAGIC_TMPFS           = 0x01021994
	S_MAGIC_TRACEFS         = 0x74726163
	S_MAGIC_UBIFS           = 0x24051905
	S_MAGIC_UDF             = 0x15013346
	S_MAGIC_UFS             = 0x00011954
	S_MAGIC_UFS_BYTESWAPPED = 0x54190100
	S_MAGIC_USBDEVFS        = 0x9FA2
	S_MAGIC_V9FS            = 0x01021997
	S_MAGIC_VMHGFS          = 0xBACBACBC
	S_MAGIC_VXFS            = 0xA501FCF5
	S_MAGIC_VZFS            = 0x565A4653
	S_MAGIC_WSLFS           = 0x53464846
	S_MAGIC_XENFS           = 0xABBA1974
	S_MAGIC_XENIX           = 0x012FF7B4
	S_MAGIC_XFS             = 0x58465342
	S_MAGIC_XIAFS           = 0x012FD16D
	S_MAGIC_ZFS             = 0x2FC12FC1
	S_MAGIC_ZSMALLOC        = 0x58295829
)
