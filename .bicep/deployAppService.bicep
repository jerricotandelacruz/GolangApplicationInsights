param appServicePlanName string

param location string = resourceGroup().location

param projectName string

param imageName string

@secure()
param containerServer string

@secure()
param appServiceSettings object

@allowed([
  'F1'
  'B1'
  'P1v2'
  'P2v2'
  'P3v2'
  'P1V3'
  'P2V3'
  'P3V3'
])
param sku string = 'P1v2'

resource appServiceLogAnalytics 'Microsoft.OperationalInsights/workspaces@2022-10-01' = {
  name: projectName
  location: location
  properties: {
    publicNetworkAccessForIngestion: 'Enabled'
    publicNetworkAccessForQuery: 'Enabled'
    sku: {
      name: 'Standard'
    }
  }
}

resource appServiceApplicationInsights 'Microsoft.Insights/components@2020-02-02' = {
  name: projectName
  location: location
  kind: 'web'
  properties: {
    Application_Type: 'web'
    RetentionInDays: 90
    WorkspaceResourceId: appServiceLogAnalytics.id
    IngestionMode: 'LogAnalytics'
    publicNetworkAccessForIngestion: 'Enabled'
    publicNetworkAccessForQuery: 'Enabled'
  }
}

resource appServicePlan 'Microsoft.Web/serverfarms@2020-06-01' = {
  name: appServicePlanName
  location: location
  properties: {
    reserved: true
  }
  sku: {
    name: sku
  }
  kind: 'linux'
}

var appServiceAppSettings = [for item in items(appServiceSettings): {
  name: item.key
  value: item.value
}]

resource appService 'Microsoft.Web/sites@2022-03-01' = {
  name: projectName
  location: location
  properties: {
    serverFarmId: appServicePlan.id
    siteConfig: {
      appSettings: union(appServiceAppSettings, [{
          name: 'APPINSIGHTS_INSTRUMENTATIONKEY'
          value: appServiceApplicationInsights.properties.InstrumentationKey
        }])
      linuxFxVersion: 'DOCKER|${containerServer}/${imageName}'
    }
  }
}
