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
			// Some padding on the right of nodes
			if (data.side === 'right') return sourceX - 5;

			const seedNode = useInternalNode(source);
			// Fall back value seed node width
			const diff = seedNode.current?.measured.width || 100;
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

	const style = data?.isDashed ? 'stroke-dasharray: 5;' : '';
</script>

<BaseEdge class={[data?.isDashed && 'opacity-40']} {style} {id} {path} />
