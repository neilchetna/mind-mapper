class FlowManager {
	openCreateNewNode = $state<boolean>(false);

	destroy() {
		this.openCreateNewNode = false;
	}
}
export const flowState = new FlowManager();
