import Configuration from "../../models/configuration/configuration";

class ConfigProvider {
// SINGLETON ---------------------
  private constructor() { }
  private _instance: ConfigProvider|null = null

  public get instance() : ConfigProvider {
    if (this._instance == null) {
      this._instance = new ConfigProvider();
    }

    return this._instance;
  }

// METHODS ---------------------
  async getConfig(): Promise<Configuration> {
    const result = await (await fetch("/api/config")).json();

    return Configuration.fromJSON(result);
  }
  
  async setConfig(config: Configuration): Promise<void> {
    await fetch("/api/config", 
      {
        method:"POST",
        body:JSON.stringify(config.asJson)
      }
    )
  }
}

export default ConfigProvider
