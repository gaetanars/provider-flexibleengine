package identity

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("flexibleengine_identity_agency_v3", func(r *config.Resource) {
		r.Kind = "Agency"
	})
	p.AddResourceConfigurator("flexibleengine_identity_group_membership_v3", func(r *config.Resource) {
		r.References["group"] = config.Reference{
			Type: "Group",
		}
		r.Kind = "GroupMembership"
	})
	p.AddResourceConfigurator("flexibleengine_identity_group_v3", func(r *config.Resource) {
		r.Kind = "Group"
	})
	p.AddResourceConfigurator("flexibleengine_identity_project_v3", func(r *config.Resource) {
		r.Kind = "Project"
	})
	p.AddResourceConfigurator("flexibleengine_identity_role_assignment_v3", func(r *config.Resource) {
		r.References["role_id"] = config.Reference{
			Type: "Role",
		}
		r.References["group_id"] = config.Reference{
			Type: "Group",
		}
		r.Kind = "RoleAssignment"
	})
	p.AddResourceConfigurator("flexibleengine_identity_role_v3", func(r *config.Resource) {
		r.Kind = "Role"
	})
	p.AddResourceConfigurator("flexibleengine_identity_user_v3", func(r *config.Resource) {
		r.Kind = "User"
	})
}
