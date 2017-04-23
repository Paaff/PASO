(function() {
  freeboard.loadDatasourcePlugin({
    "type_name" : "discoveredBluetooth_datasource_plugin",
    "display_name" : "Discovered Bluetooth Clients Plugin",
    "description" : "This datasource will provide who has been discovered by bluetooth detection.",
    "settings" : [
      {
        "name" : "setting",
        "display_name" : "Placeholder setting",
        "type" : "option",
        "options" : [
          {
            "name" : "First Option",
            "value": "First"
          },
          {
            "name" : "Second Option",
            "value": "Second"
          },
          {
            "name": "Third Option",
            "value": "Third"
          }
        ]
      },
      {
        "name"         : "refresh_time",
				"display_name" : "Refresh Time",
				"type"         : "text",
				"description"  : "In milliseconds",
				"default_value": 5000
      }
    ],
    newInstance : function(settings, newInstanceCallback, updateCallback) {
      newInstanceCallback(new myDatasourcePlugin(settings, updateCallback));
    }


  });

  var myDatasourcePlugin = function(settings, updateCallback) {
    var self = this;
    var currentSettings = settings;
    var counter = 0;

    /* This is some function where I'll get my data from somewhere */
		function getData()
		{


			var newData = counter++; // Just putting some sample data in for fun.

			/* Get my data from somewhere and populate newData with it... Probably a JSON API or something. */
			/* ... */
      updateCallback(newData);
  }

  var refreshTimer;

  function createRefreshTimer(interval) {
    if(refreshTimer)
			{
				clearInterval(refreshTimer);
			}

			refreshTimer = setInterval(function()
			{
        getData();
      }, interval);
  }

  self.onSettingsChanged = function(newSettings) {
    currentSettings = newSettings;
  };

  self.updateNow = function() {
    getData();
  };

  self.onDispose = function() {
    clearInterval(refreshTimer);
    refreshTimer = undefined;
  };

  createRefreshTimer(currentSettings.refresh_time);
};

freeboard.loadWidgetPlugin({
  "type_name" : "clients_discovered_widget_plugin",
  "display_name" : "Clients Discovered Widget",
  "description" : "Will show the clients discovered.",
  "fill_size" : false,
  "settings" : [
    {
      "name" : "client_bt_address",
      "display_name" : "Bluetooth Address",
      "type" : "calculated"
    },
    {
      "name"        : "size",
				"display_name": "Size",
				"type"        : "option",
				"options"     : [
					{
						"name" : "Regular",
						"value": "regular"
					},
					{
						"name" : "Big",
						"value": "big"
					}
				]

    }

  ],

  newInstance : function(settings, newInstanceCallback) {
    newInstanceCallback(new myWidgetPlugin(settings));
  }
});

var myWidgetPlugin = function(settings) {
  var self = this;
  var currentSettings = settings;
  var myTextElement = $("<span></span>");
  self.render = function(containerElement) {
    $(containerElement).append(myTextElement);
  };

  self.getHeight = function() {
    if(currentSettings.size == "big")
			{
				return 2;
			}
			else
			{
				return 1;
			}
  };


  self.onSettingsChanged = function(newSettings) {
  currentSettings = newSettings;
};

  self.onCalculatedValueChanged = function(settingName, newValue) {
    if(settingName == "client_bt_address") {
      $(myTextElement).html(newValue);
    }
  };

  self.onDispose = function() {
  };
};


}());
