// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-2020 Datadog, Inc.

// +build linux

package probe

import (
	"github.com/DataDog/datadog-agent/pkg/security/secl/eval"
)

// KProbe describes a Linux Kprobe
type KProbe struct {
	Name      string
	EntryFunc string
	ExitFunc  string
}

// HookPoint represents
type HookPoint struct {
	Name            string
	KProbes         []*KProbe
	Tracepoint      string
	Optional        bool
	EventTypes      []eval.EventType
	OnNewApprovers  onApproversFnc
	OnNewDiscarders onDiscarderFnc
	PolicyTable     string
}

var allHookPoints = []*HookPoint{
	{
		Name: "security_inode_setattr",
		KProbes: []*KProbe{{
			EntryFunc: "kprobe/security_inode_setattr",
		}},
		EventTypes: []eval.EventType{"chmod", "chown", "utimes"},
	},
	{
		Name:       "sys_chmod",
		KProbes:    syscallKprobe("chmod"),
		EventTypes: []eval.EventType{"chmod"},
	},
	{
		Name:       "sys_fchmod",
		KProbes:    syscallKprobe("fchmod"),
		EventTypes: []eval.EventType{"chmod"},
	},
	{
		Name:       "sys_fchmodat",
		KProbes:    syscallKprobe("fchmodat"),
		EventTypes: []eval.EventType{"chmod"},
	},
	{
		Name:       "sys_chown",
		KProbes:    syscallKprobe("chown"),
		EventTypes: []eval.EventType{"chown"},
	},
	{
		Name:       "sys_fchown",
		KProbes:    syscallKprobe("fchown"),
		EventTypes: []eval.EventType{"chown"},
	},
	{
		Name:       "sys_fchownat",
		KProbes:    syscallKprobe("fchownat"),
		EventTypes: []eval.EventType{"chown"},
	},
	{
		Name:       "sys_lchown",
		KProbes:    syscallKprobe("lchown"),
		EventTypes: []eval.EventType{"chown"},
	},
	{
		Name: "mnt_want_write",
		KProbes: []*KProbe{{
			EntryFunc: "kprobe/mnt_want_write",
		}},
		EventTypes: []eval.EventType{"utimes", "chmod", "chown", "rmdir", "unlink", "rename"},
	},
	{
		Name: "mnt_want_write_file",
		KProbes: []*KProbe{{
			EntryFunc: "kprobe/mnt_want_write_file",
		}},
		EventTypes: []eval.EventType{"chown"},
	},
	{
		Name:       "sys_utime",
		KProbes:    syscallKprobe("utime"),
		EventTypes: []eval.EventType{"utimes"},
	},
	{
		Name:       "sys_utimes",
		KProbes:    syscallKprobe("utimes"),
		EventTypes: []eval.EventType{"utimes"},
	},
	{
		Name:       "sys_utimensat",
		KProbes:    syscallKprobe("utimensat"),
		EventTypes: []eval.EventType{"utimes"},
	},
	{
		Name:       "sys_futimesat",
		KProbes:    syscallKprobe("futimesat"),
		EventTypes: []eval.EventType{"utimes"},
	},
	{
		Name: "vfs_mkdir",
		KProbes: []*KProbe{{
			EntryFunc: "kprobe/vfs_mkdir",
		}},
		EventTypes: []eval.EventType{"mkdir"},
	},
	{
		Name: "filename_create",
		KProbes: []*KProbe{{
			EntryFunc: "kprobe/filename_create",
		}},
		EventTypes: []eval.EventType{"mkdir", "link"},
	},
	{
		Name:       "sys_mkdir",
		KProbes:    syscallKprobe("mkdir"),
		EventTypes: []eval.EventType{"mkdir"},
	},
	{
		Name:       "sys_mkdirat",
		KProbes:    syscallKprobe("mkdirat"),
		EventTypes: []eval.EventType{"mkdir"},
	},
	{
		Name: "vfs_rmdir",
		KProbes: []*KProbe{{
			EntryFunc: "kprobe/vfs_rmdir",
		}},
		EventTypes: []eval.EventType{"rmdir", "unlink"},
	},
	{
		Name:       "sys_rmdir",
		KProbes:    syscallKprobe("rmdir"),
		EventTypes: []eval.EventType{"rmdir"},
	},
	{
		Name: "vfs_rename",
		KProbes: []*KProbe{{
			EntryFunc: "kprobe/vfs_rename",
		}},
		EventTypes: []eval.EventType{"rename"},
	},
	{
		Name:       "sys_rename",
		KProbes:    syscallKprobe("rename"),
		EventTypes: []eval.EventType{"rename"},
	},
	{
		Name:       "sys_renameat",
		KProbes:    syscallKprobe("renameat"),
		EventTypes: []eval.EventType{"rename"},
	},
	{
		Name:       "sys_renameat2",
		KProbes:    syscallKprobe("renameat2"),
		EventTypes: []eval.EventType{"rename"},
	},
	{
		Name: "vfs_link",
		KProbes: []*KProbe{{
			EntryFunc: "kprobe/vfs_link",
		}},
		EventTypes: []eval.EventType{"link"},
	},
	{
		Name:       "sys_link",
		KProbes:    syscallKprobe("link"),
		EventTypes: []eval.EventType{"link"},
	},
	{
		Name:       "sys_linkat",
		KProbes:    syscallKprobe("linkat"),
		EventTypes: []eval.EventType{"link"},
	},
}

func init() {
	allHookPoints = append(allHookPoints, openHookPoints...)
	allHookPoints = append(allHookPoints, mountHookPoints...)
	allHookPoints = append(allHookPoints, execHookPoints...)
	allHookPoints = append(allHookPoints, UnlinkHookPoints...)
}