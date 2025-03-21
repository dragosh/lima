package limayaml

import (
	"net"

	"github.com/opencontainers/go-digest"
)

type LimaYAML struct {
	Arch              Arch              `yaml:"arch,omitempty" json:"arch,omitempty"`
	Images            []File            `yaml:"images" json:"images"` // REQUIRED
	CPUs              int               `yaml:"cpus,omitempty" json:"cpus,omitempty"`
	Memory            string            `yaml:"memory,omitempty" json:"memory,omitempty"` // go-units.RAMInBytes
	Disk              string            `yaml:"disk,omitempty" json:"disk,omitempty"`     // go-units.RAMInBytes
	Mounts            []Mount           `yaml:"mounts,omitempty" json:"mounts,omitempty"`
	SSH               SSH               `yaml:"ssh,omitempty" json:"ssh,omitempty"` // REQUIRED (FIXME)
	Firmware          Firmware          `yaml:"firmware,omitempty" json:"firmware,omitempty"`
	Video             Video             `yaml:"video,omitempty" json:"video,omitempty"`
	Provision         []Provision       `yaml:"provision,omitempty" json:"provision,omitempty"`
	Containerd        Containerd        `yaml:"containerd,omitempty" json:"containerd,omitempty"`
	Probes            []Probe           `yaml:"probes,omitempty" json:"probes,omitempty"`
	PortForwards      []PortForward     `yaml:"portForwards,omitempty" json:"portForwards,omitempty"`
	Networks          []Network         `yaml:"networks,omitempty" json:"networks,omitempty"`
	Network           NetworkDeprecated `yaml:"network,omitempty" json:"network,omitempty"` // DEPRECATED, use `networks` instead
	Env               map[string]string `yaml:"env,omitempty" json:"env,omitempty"`
	DNS               []net.IP          `yaml:"dns,omitempty" json:"dns,omitempty"`
	UseHostResolver   *bool             `yaml:"useHostResolver,omitempty" json:"useHostResolver,omitempty"`
	PropagateProxyEnv *bool             `yaml:"propagateProxyEnv,omitempty" json:"propagateProxyEnv,omitempty"`
}

type Arch = string

const (
	X8664   Arch = "x86_64"
	AARCH64 Arch = "aarch64"
)

type File struct {
	Location string        `yaml:"location" json:"location"` // REQUIRED
	Arch     Arch          `yaml:"arch,omitempty" json:"arch,omitempty"`
	Digest   digest.Digest `yaml:"digest,omitempty" json:"digest,omitempty"`
}

type Mount struct {
	Location string `yaml:"location" json:"location"` // REQUIRED
	Writable bool   `yaml:"writable,omitempty" json:"writable,omitempty"`
}

type SSH struct {
	LocalPort int `yaml:"localPort,omitempty" json:"localPort,omitempty"`

	// LoadDotSSHPubKeys loads ~/.ssh/*.pub in addition to $LIMA_HOME/_config/user.pub .
	LoadDotSSHPubKeys *bool `yaml:"loadDotSSHPubKeys,omitempty" json:"loadDotSSHPubKeys,omitempty"` // default: true
	ForwardAgent      *bool `yaml:"forwardAgent,omitempty" json:"forwardAgent,omitempty"`           // default: false
}

type Firmware struct {
	// LegacyBIOS disables UEFI if set.
	// LegacyBIOS is ignored for aarch64.
	LegacyBIOS bool `yaml:"legacyBIOS,omitempty" json:"legacyBIOS,omitempty"`
}

type Video struct {
	// Display is a QEMU display string
	Display string `yaml:"display,omitempty" json:"display,omitempty"`
}

type ProvisionMode = string

const (
	ProvisionModeSystem ProvisionMode = "system"
	ProvisionModeUser   ProvisionMode = "user"
)

type Provision struct {
	Mode   ProvisionMode `yaml:"mode" json:"mode"` // default: "system"
	Script string        `yaml:"script" json:"script"`
}

type Containerd struct {
	System   *bool  `yaml:"system,omitempty" json:"system,omitempty"`     // default: false
	User     *bool  `yaml:"user,omitempty" json:"user,omitempty"`         // default: true
	Archives []File `yaml:"archives,omitempty" json:"archives,omitempty"` // default: see defaultContainerdArchives
}

type ProbeMode = string

const (
	ProbeModeReadiness ProbeMode = "readiness"
)

type Probe struct {
	Mode        ProbeMode // default: "readiness"
	Description string
	Script      string
	Hint        string
}

type Proto = string

const (
	TCP Proto = "tcp"
)

type PortForward struct {
	GuestIP        net.IP `yaml:"guestIP,omitempty" json:"guestIP,omitempty"`
	GuestPort      int    `yaml:"guestPort,omitempty" json:"guestPort,omitempty"`
	GuestPortRange [2]int `yaml:"guestPortRange,omitempty" json:"guestPortRange,omitempty"`
	HostIP         net.IP `yaml:"hostIP,omitempty" json:"hostIP,omitempty"`
	HostPort       int    `yaml:"hostPort,omitempty" json:"hostPort,omitempty"`
	HostPortRange  [2]int `yaml:"hostPortRange,omitempty" json:"hostPortRange,omitempty"`
	Proto          Proto  `yaml:"proto,omitempty" json:"proto,omitempty"`
	Ignore         bool   `yaml:"ignore,omitempty" json:"ignore,omitempty"`
}

type Network struct {
	// `Lima` and `VNL` are mutually exclusive; exactly one is required
	Lima string `yaml:"lima,omitempty" json:"lima,omitempty"`
	// VNL is a Virtual Network Locator (https://github.com/rd235/vdeplug4/commit/089984200f447abb0e825eb45548b781ba1ebccd).
	// On macOS, only VDE2-compatible form (optionally with vde:// prefix) is supported.
	VNL        string `yaml:"vnl,omitempty" json:"vnl,omitempty"`
	SwitchPort uint16 `yaml:"switchPort,omitempty" json:"switchPort,omitempty"` // VDE Switch port, not TCP/UDP port (only used by VDE networking)
	MACAddress string `yaml:"macAddress,omitempty" json:"macAddress,omitempty"`
	Interface  string `yaml:"interface,omitempty" json:"interface,omitempty"`
}

// DEPRECATED types below

// Types have been renamed to turn all references to the old names into compiler errors,
// and to avoid accidental usage in new code.

type NetworkDeprecated struct {
	VDEDeprecated []VDEDeprecated `yaml:"vde,omitempty" json:"vde,omitempty"`
	// migrate will be true when `network.VDE` has been copied to `networks` by FillDefaults()
	migrated bool
}

type VDEDeprecated struct {
	VNL        string `yaml:"vnl,omitempty" json:"vnl,omitempty"`
	SwitchPort uint16 `yaml:"switchPort,omitempty" json:"switchPort,omitempty"` // VDE Switch port, not TCP/UDP port
	MACAddress string `yaml:"macAddress,omitempty" json:"macAddress,omitempty"`
	Name       string `yaml:"name,omitempty" json:"name,omitempty"`
}
