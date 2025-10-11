package boot

import (
	"github.com/EQEmuTools/spirerere/internal/assets"
	"github.com/EQEmuTools/spirerere/internal/auditlog"
	"github.com/EQEmuTools/spirerere/internal/backup"
	"github.com/EQEmuTools/spirerere/internal/clientfiles"
	"github.com/EQEmuTools/spirerere/internal/connection"
	"github.com/EQEmuTools/spirerere/internal/desktop"
	"github.com/EQEmuTools/spirerere/internal/eqemuanalytics"
	"github.com/EQEmuTools/spirerere/internal/eqemuchangelog"
	"github.com/EQEmuTools/spirerere/internal/eqemuloginserver"
	"github.com/EQEmuTools/spirerere/internal/eqemuserver"
	"github.com/EQEmuTools/spirerere/internal/eqemuserverconfig"
	"github.com/EQEmuTools/spirerere/internal/github"
	"github.com/EQEmuTools/spirerere/internal/influx"
	"github.com/EQEmuTools/spirerere/internal/pathmgmt"
	"github.com/EQEmuTools/spirerere/internal/permissions"
	"github.com/EQEmuTools/spirerere/internal/questapi"
	"github.com/EQEmuTools/spirerere/internal/spire"
	"github.com/EQEmuTools/spirerere/internal/telnet"
	"github.com/EQEmuTools/spirerere/internal/unzip"
	"github.com/EQEmuTools/spirerere/internal/user"
	"github.com/EQEmuTools/spirerere/internal/websocket"
	pluralize "github.com/gertd/go-pluralize"
	"github.com/google/wire"
)

var serviceSet = wire.NewSet(
	influx.NewClient,
	connection.NewCreate,
	connection.NewCheck,
	github.NewGithubSourceDownloader,
	questapi.NewParseService,
	questapi.NewExamplesGithubSourcer,
	desktop.NewWebBoot,
	clientfiles.NewExporter,
	clientfiles.NewImporter,
	eqemuserverconfig.NewConfig,
	eqemuloginserver.NewConfig,
	pathmgmt.NewPathManagement,
	permissions.NewService,
	pluralize.NewClient,
	auditlog.NewUserEvent,
	assets.NewSpireAssets,
	eqemuchangelog.NewChangelog,
	eqemuanalytics.NewReleases,
	user.NewUser,
	spire.NewSettings,
	spire.NewInit,
	telnet.NewClient,
	eqemuserver.NewClient,
	backup.NewMysql,
	websocket.NewHandler,
	eqemuserver.NewUpdater,
	eqemuserver.NewLauncher,
	eqemuserver.NewQuestHotReloadWatcher,
	unzip.NewUnzipper,
	websocket.NewClientManager,
	eqemuserver.NewCrashLogWatcher,
)
