package push

import (
	"code.cloudfoundry.org/cli/integration/helpers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
	. "github.com/onsi/gomega/gexec"
)

var _ = Describe("push with hostname", func() {
	Context("when the default domain is a shared domain", func() {
		DescribeTable("creates and binds the route as neccessary",
			func(existingRoute bool, boundRoute bool, setup func(appName string, dir string) *Session) {
				appName := helpers.NewAppName()

				if existingRoute {
					session := helpers.CF("create-route", space, defaultSharedDomain(), "-n", appName)
					Eventually(session).Should(Exit(0))
				}

				if boundRoute {
					helpers.WithHelloWorldApp(func(dir string) {
						// TODO: Add --no-start
						session := helpers.CustomCF(helpers.CFEnv{WorkingDirectory: dir}, PushCommandName, appName)
						Eventually(session).Should(Exit(0))
					})
				}

				helpers.WithHelloWorldApp(func(dir string) {
					session := setup(appName, dir)

					Eventually(session).Should(Say("routes:"))
					Eventually(session).Should(Say("(?i)%s.%s", appName, defaultSharedDomain()))
					Eventually(session).Should(Say("Configuring routes..."))

					if !existingRoute {
						Eventually(session).Should(Say("Created routes."))
					}

					if !boundRoute {
						Eventually(session).Should(Say("Bound routes."))
					}

					Eventually(session).Should(Exit(0))
				})
			},

			Entry("when the hostname is provided via the appName and route does not exist", false, false, func(appName string, dir string) *Session {
				// TODO: Add --no-start
				// TODO: Add --path
				return helpers.CustomCF(helpers.CFEnv{WorkingDirectory: dir}, PushCommandName, appName)
			}),

			Entry("when the hostname is provided via the appName and the unbound route exists", true, false, func(appName string, dir string) *Session {
				// TODO: Add --no-start
				// TODO: Add --path
				return helpers.CustomCF(helpers.CFEnv{WorkingDirectory: dir}, PushCommandName, appName)
			}),

			Entry("when the hostname is provided via the appName and the bound route exists", true, true, func(appName string, dir string) *Session {
				// TODO: Add --no-start
				// TODO: Add --path
				return helpers.CustomCF(helpers.CFEnv{WorkingDirectory: dir}, PushCommandName, appName)
			}),
		)
	})

	Context("when the default domain is private", func() {
		DescribeTable("creates and binds the route as neccessary",
			func(existingRoute bool, boundRoute bool, setup func(appName string, dir string) *Session) {
				appName := helpers.NewAppName()

				if existingRoute {
					session := helpers.CF("create-route", space, defaultSharedDomain(), "-n", appName)
					Eventually(session).Should(Exit(0))
				}

				if boundRoute {
					helpers.WithHelloWorldApp(func(dir string) {
						// TODO: Add --no-start
						session := helpers.CustomCF(helpers.CFEnv{WorkingDirectory: dir}, PushCommandName, appName)
						Eventually(session).Should(Exit(0))
					})
				}

				helpers.WithHelloWorldApp(func(dir string) {
					session := setup(appName, dir)

					Eventually(session).Should(Say("routes:"))
					Eventually(session).Should(Say("(?i)%s.%s", appName, defaultSharedDomain()))
					Eventually(session).Should(Say("Configuring routes..."))

					if !existingRoute {
						Eventually(session).Should(Say("Created routes."))
					}

					if !boundRoute {
						Eventually(session).Should(Say("Bound routes."))
					}

					Eventually(session).Should(Exit(0))
				})
			},

			Entry("when the hostname is provided via the appName and route does not exist", false, false, func(appName string, dir string) *Session {
				// TODO: Add --no-start
				// TODO: Add --path
				return helpers.CustomCF(helpers.CFEnv{WorkingDirectory: dir}, PushCommandName, appName)
			}),

			Entry("when the hostname is provided via the appName and the unbound route exists", true, false, func(appName string, dir string) *Session {
				// TODO: Add --no-start
				// TODO: Add --path
				return helpers.CustomCF(helpers.CFEnv{WorkingDirectory: dir}, PushCommandName, appName)
			}),

			Entry("when the hostname is provided via the appName and the bound route exists", true, true, func(appName string, dir string) *Session {
				// TODO: Add --no-start
				// TODO: Add --path
				return helpers.CustomCF(helpers.CFEnv{WorkingDirectory: dir}, PushCommandName, appName)
			}),
		)
	})

	// Describe("route existence", func() {
	// 	Context("when the route does not exist", func() {
	// 		It("creates and binds the route", func() {
	// 			helpers.WithHelloWorldApp(func(dir string) {
	// 				session := helpers.CustomCF(helpers.CFEnv{WorkingDirectory: dir}, PushCommandName, appName)
	// 				Eventually(session).Should(Say("Creating route %s.%s...", strings.ToLower(appName), defaultSharedDomain()))
	// 				Eventually(session).Should(Say("OK"))

	// 				Eventually(session).Should(Say("Binding %s.%s to %s...", strings.ToLower(appName), defaultSharedDomain(), appName))
	// 				Eventually(session).Should(Say("OK"))
	// 				Eventually(session).Should(Exit(0))
	// 			})
	// 		})
	// 	})

	// 	Context("when the route exists in the current space", func() {
	// 		BeforeEach(func() {
	// 			session := helpers.CF("create-route", space, defaultSharedDomain(), "-n", appName)
	// 			Eventually(session).Should(Exit(0))
	// 		})

	// 		It("should not create the route", func() {
	// 			helpers.WithHelloWorldApp(func(dir string) {
	// 				session := helpers.CustomCF(helpers.CFEnv{WorkingDirectory: dir}, PushCommandName, appName)
	// 				Consistently(session).ShouldNot(Say("Creating route %s.%s...", strings.ToLower(appName), defaultSharedDomain()))
	// 				Eventually(session).Should(Exit(0))
	// 			})
	// 		})

	// 		Context("when the route is not bound to an app of the same name", func() {
	// 			It("binds the route to the app", func() {
	// 				helpers.WithHelloWorldApp(func(dir string) {
	// 					session := helpers.CustomCF(helpers.CFEnv{WorkingDirectory: dir}, PushCommandName, appName)
	// 					Eventually(session).Should(Say("Using route %s.%s", strings.ToLower(appName), defaultSharedDomain()))
	// 					Eventually(session).Should(Say("Binding %s.%s to %s...", strings.ToLower(appName), defaultSharedDomain(), appName))
	// 					Eventually(session).Should(Say("OK"))
	// 					Eventually(session).Should(Exit(0))
	// 				})
	// 			})
	// 		})

	// 		Context("when the route is already bound to the application", func() {
	// 			BeforeEach(func() {
	// 				helpers.WithHelloWorldApp(func(dir string) {
	// 					session := helpers.CustomCF(helpers.CFEnv{WorkingDirectory: dir}, PushCommandName, appName)
	// 					Eventually(session).Should(Exit(0))
	// 				})
	// 			})

	// 			It("does not rebind the route", func() {
	// 				helpers.WithHelloWorldApp(func(dir string) {
	// 					session := helpers.CustomCF(helpers.CFEnv{WorkingDirectory: dir}, PushCommandName, appName)
	// 					Consistently(session).ShouldNot(Say("Binding %s.%s to %s...", strings.ToLower(appName), defaultSharedDomain(), appName))
	// 					Eventually(session).Should(Exit(0))
	// 				})
	// 			})
	// 		})
	// 	})

	// 	Context("when the route exists in a different space", func() {
	// 		BeforeEach(func() {
	// 			otherSpace := helpers.NewSpaceName()
	// 			Eventually(helpers.CF("create-space", otherSpace)).Should(Exit(0))
	// 			Eventually(helpers.CF("create-route", otherSpace, defaultSharedDomain(), "-n", appName)).Should(Exit(0))
	// 		})

	// 		It("errors", func() {
	// 			helpers.WithHelloWorldApp(func(dir string) {
	// 				session := helpers.CustomCF(helpers.CFEnv{WorkingDirectory: dir}, PushCommandName, appName)
	// 				Eventually(session).Should(Say("Using route %s.%s", strings.ToLower(appName), defaultSharedDomain()))
	// 				Eventually(session).Should(Say("Binding %s.%s to %s...", strings.ToLower(appName), defaultSharedDomain(), appName))
	// 				Eventually(session).Should(Say("FAILED"))
	// 				Eventually(session.Err).Should(Say("The route %s.%s is already in use.", appName, defaultSharedDomain()))
	// 				Eventually(session.Err).Should(Say("TIP: Change the hostname with -n HOSTNAME or use --random-route to generate a new route and then push again."))
	// 				Eventually(session).Should(Exit(1))
	// 			})
	// 		})
	// 	})
	// })
})
