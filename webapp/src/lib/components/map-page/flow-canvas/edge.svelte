<script lang="ts">
	import {
		BaseEdge,
		getBezierPath,
		Position,
		useInternalNode,
		type EdgeProps
	} from '@xyflow/svelte';

	const { id, sourceX, targetX, sourceY, targetY, data, source }: EdgeProps = $props();

	function getSourcePosition(): Position.Left | Position.Right {
		return data?.side === 'left' ? Position.Left : Position.Right;
	}

	function getTargetPosition(): Position.Left | Position.Right {
		return data?.side === 'left' ? Position.Right : Position.Left;
	}

	function getSourceX() {
		if (data?.isSeedEdge) {
			if (data.side === 'right') return sourceX - 5; // Some padding on the right of nodes

			const seedNode = useInternalNode(source);
			const diff = seedNode.current?.measured.width || 100; // Fall back value seed node width
			return sourceX - diff;
		}
		return sourceX;
	}

	let [path] = $derived(
		getBezierPath({
			sourceX: getSourceX(),
			sourceY,
			sourcePosition: getSourcePosition(),
			targetPosition: getTargetPosition(),
			targetX,
			targetY
		})
	);
</script>

<BaseEdge {id} {path} />
