package supervisord

import (
	"fmt"
)

// WithGRPCAddress sets the containerd address.
func WithGRPCAddress(addr string) Opt {
	return func(d *Daemon) error {
		if addr == "" {
			return fmt.Errorf("grpc address should not be empty")
		}

		d.cfg.GRPC.Address = addr
		return nil
	}
}

// WithLogLevel sets the log level of containerd.
func WithLogLevel(level string) Opt {
	return func(d *Daemon) error {
		d.cfg.Debug.Level = level
		return nil
	}
}

// WithOOMScore sets the OOMScore for containerd.
func WithOOMScore(score int) Opt {
	return func(d *Daemon) error {
		if score > 1000 || score < -1000 {
			return fmt.Errorf("oom-score range should be [-1000, 1000]")
		}

		d.cfg.OOMScore = score
		return nil
	}
}

// WithContainerdBinary sets the binary name or path of containerd.
func WithContainerdBinary(nameOrPath string) Opt {
	return func(d *Daemon) error {
		d.binaryName = nameOrPath
		return nil
	}
}

// WithSnapshotterConfig passes down snapshotter config to containerd
func WithSnapshotterConfig(s string, c interface{}) Opt {
	return func(d *Daemon) error {
		if d.cfg.Plugins == nil {
			d.cfg.Plugins = map[string]interface{}{}
		}

		d.cfg.Plugins[s] = c
		return nil
	}
}

// WithV1RuntimeConfig set v1 runtime config in containerd
func WithV1RuntimeConfig(shim string) Opt {
	return func(d *Daemon) error {

		if d.cfg.Plugins == nil {
			d.cfg.Plugins = map[string]interface{}{}
		}

		v1RuntimeCfg := V1RuntimeConfig{}

		// always enable debug option
		v1RuntimeCfg.ShimDebug = true

		if shim != "" {
			v1RuntimeCfg.Shim = shim
		}

		d.cfg.Plugins["linux"] = v1RuntimeCfg

		return nil
	}
}

// WithProxyPluginConfig passes down proxy plugin config to containerd
func WithProxyPluginConfig(s string, c map[string]string) Opt {
	return func(d *Daemon) error {
		if d.cfg.ProxyPlugins == nil {
			d.cfg.ProxyPlugins = map[string]ProxyPlugin{}
		}

		d.cfg.ProxyPlugins[s] = ProxyPlugin{
			Type:    c["type"],
			Address: c["address"],
		}
		return nil
	}
}
