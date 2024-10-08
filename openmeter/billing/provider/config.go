package provider

import "fmt"

// Type specifies the provider used for billing
type Type string

const (
	// TypeOpenMeter specifies the OpenMeter billing provider, which is a dummy billing provider mostly useful for testing and
	// initial OpenMeter assessment
	TypeOpenMeter Type = "openmeter"
	// TypeStripe specifies the Stripe billing provider, which is a real billing provider that can be used in production
	TypeStripe Type = "stripe"
)

func (t Type) Values() []string {
	return []string{
		string(TypeOpenMeter),
		string(TypeStripe),
	}
}

type Meta struct {
	Type Type `json:"type"`
}

type Configuration struct {
	Meta

	OpenMeter OpenMeterConfig `json:"openMeter"`
	Stripe    StripeConfig    `json:"stripe"`
}

func (c *Configuration) Validate() error {
	switch c.Type {
	case TypeOpenMeter:
		if err := c.OpenMeter.Validate(); err != nil {
			return fmt.Errorf("failed to validate openmeter configuration: %w", err)
		}

	case TypeStripe:
		if err := c.Stripe.Validate(); err != nil {
			return fmt.Errorf("failed to validate stripe configuration: %w", err)
		}

	default:
		return fmt.Errorf("unknown backend type: %s", c.Type)
	}

	return nil
}

type OpenMeterConfig struct{}

func (c *OpenMeterConfig) Validate() error {
	return nil
}

type StripeConfig struct{}

func (c *StripeConfig) Validate() error {
	return nil
}
